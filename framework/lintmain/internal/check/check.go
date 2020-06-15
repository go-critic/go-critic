package check

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/go-critic/go-critic/framework/linter"
	"github.com/go-critic/go-critic/framework/lintmain/internal/hotload"
	"github.com/go-toolsmith/pkgload"
	"github.com/logrusorgru/aurora"
	"golang.org/x/tools/go/packages"
)

// Main implements sub-command entry point.
func Main() {
	var p program
	p.infoList = linter.GetCheckersInfo()

	steps := []struct {
		name string
		fn   func() error
	}{
		{"load plugin", p.loadPlugin},
		{"bind checker params", p.bindCheckerParams},
		{"bind default enabled list", p.bindDefaultEnabledList},
		{"parse args", p.parseArgs},
		{"assign checker params", p.assignCheckerParams},
		{"load program", p.loadProgram},
		{"init checkers", p.initCheckers},
		{"run checkers", p.runCheckers},
		{"exit if found issues", p.exit},
	}

	for _, step := range steps {
		if err := step.fn(); err != nil {
			log.Fatalf("%s: %v", step.name, err)
		}
	}
}

type program struct {
	ctx *linter.Context

	fset *token.FileSet

	loadedPackages []*packages.Package

	infoList []*linter.CheckerInfo

	checkers []*linter.Checker

	packages []string

	foundIssues bool

	checkerParams boundCheckerParams

	filters struct {
		enableAll       bool
		enable          []string
		disable         []string
		defaultCheckers []string
	}

	workDir string
	gopath  string
	goroot  string

	exitCode           int
	checkTests         bool
	checkGenerated     bool
	shorterErrLocation bool
	coloredOutput      bool
	verbose            bool
}

func (p *program) exit() error {
	if p.foundIssues {
		os.Exit(p.exitCode)
	}
	return nil
}

func (p *program) runCheckers() error {
	for _, pkg := range p.loadedPackages {
		if p.verbose {
			log.Printf("\tdebug: checking %q package (%d files)",
				pkg.String(), len(pkg.Syntax))
		}
		p.checkPackage(pkg)
	}

	return nil
}

func (p *program) checkPackage(pkg *packages.Package) {
	p.ctx.SetPackageInfo(pkg.TypesInfo, pkg.Types)
	for _, f := range pkg.Syntax {
		filename := p.getFilename(f)
		if !p.checkTests && strings.HasSuffix(filename, "_test.go") {
			continue
		}
		if !p.checkGenerated && p.isGenerated(f) {
			continue
		}
		p.ctx.SetFileInfo(filename, f)
		p.checkFile(f)
	}
}

func (p *program) checkFile(f *ast.File) {
	warnings := make([][]linter.Warning, len(p.checkers))

	var wg sync.WaitGroup
	wg.Add(len(p.checkers))
	for i, c := range p.checkers {
		// All checkers are expected to use *lint.Context
		// as read-only structure, so no copying is required.
		go func(i int, c *linter.Checker) {
			defer func() {
				wg.Done()
				// Checker signals unexpected error with panic(error).
				r := recover()
				if r == nil {
					return // There were no panic
				}
				if err, ok := r.(error); ok {
					log.Printf("%s: error: %v\n", c.Info.Name, err)
					panic(err)
				} else {
					// Some other kind of run-time panic.
					// Undo the recover and resume panic.
					panic(r)
				}
			}()

			for _, warn := range c.Check(f) {
				warnings[i] = append(warnings[i], warn)
			}
		}(i, c)
	}
	wg.Wait()

	for i, c := range p.checkers {
		for _, warn := range warnings[i] {
			p.foundIssues = true
			loc := p.ctx.FileSet.Position(warn.Node.Pos()).String()
			if p.shorterErrLocation {
				loc = p.shortenLocation(loc)
			}
			printWarning(p, c.Info.Name, loc, warn.Text)
		}
	}

}

func (p *program) initCheckers() error {
	parseKeys := func(keys []string, byName, byTag map[string]bool) {
		for _, key := range keys {
			if strings.HasPrefix(key, "#") {
				byTag[key[len("#"):]] = true
			} else {
				byName[key] = true
			}
		}
	}

	enabledByName := make(map[string]bool)
	enabledTags := make(map[string]bool)
	parseKeys(p.filters.enable, enabledByName, enabledTags)
	disabledByName := make(map[string]bool)
	disabledTags := make(map[string]bool)
	parseKeys(p.filters.disable, disabledByName, disabledTags)

	enabledByTag := func(info *linter.CheckerInfo) bool {
		for _, tag := range info.Tags {
			if enabledTags[tag] {
				return true
			}
		}
		return false
	}
	disabledByTag := func(info *linter.CheckerInfo) string {
		for _, tag := range info.Tags {
			if disabledTags[tag] {
				return tag
			}
		}
		return ""
	}

	for _, info := range p.infoList {
		enabled := p.filters.enableAll ||
			enabledByName[info.Name] ||
			enabledByTag(info)
		notice := ""

		switch {
		case !enabled:
			notice = "not enabled by name or tag (-enable)"
		case disabledByName[info.Name]:
			enabled = false
			notice = "disabled by name (-disable)"
		default:
			if tag := disabledByTag(info); tag != "" {
				enabled = false
				notice = fmt.Sprintf("disabled by %q tag (-disable)", tag)
			}
		}

		if p.verbose && !enabled {
			log.Printf("\tdebug: %s: %s", info.Name, notice)
		}
		if enabled {
			p.checkers = append(p.checkers, linter.NewChecker(p.ctx, info))
		}
	}
	if p.verbose {
		for _, c := range p.checkers {
			log.Printf("\tdebug: %s is enabled", c.Info.Name)
		}
	}

	if len(p.checkers) == 0 {
		return errors.New("empty checkers set selected")
	}
	return nil
}

func (p *program) loadProgram() error {
	sizes := types.SizesFor("gc", runtime.GOARCH)
	if sizes == nil {
		return fmt.Errorf("can't find sizes info for %s", runtime.GOARCH)
	}

	p.fset = token.NewFileSet()
	cfg := packages.Config{
		Mode:  packages.LoadSyntax,
		Tests: true,
		Fset:  p.fset,
	}
	pkgs, err := loadPackages(&cfg, p.packages)
	if err != nil {
		log.Fatalf("load packages: %v", err)
	}
	sort.SliceStable(pkgs, func(i, j int) bool {
		return pkgs[i].PkgPath < pkgs[j].PkgPath
	})

	p.loadedPackages = pkgs
	p.ctx = linter.NewContext(p.fset, sizes)

	return nil
}

func (p *program) loadPlugin() error {
	const pluginFilename = "gocritic-plugin.so"
	if _, err := os.Stat(pluginFilename); os.IsNotExist(err) {
		return nil
	}
	infoList, err := hotload.CheckersFromDylib(p.infoList, pluginFilename)
	p.infoList = infoList
	return err
}

type boundCheckerParams struct {
	ints    map[string]*int
	bools   map[string]*bool
	strings map[string]*string
}

// bindCheckerParams registers command-line flags for every checker parameter.
func (p *program) bindCheckerParams() error {
	intParams := make(map[string]*int)
	boolParams := make(map[string]*bool)
	stringParams := make(map[string]*string)

	for _, info := range p.infoList {
		for pname, param := range info.Params {
			key := p.checkerParamKey(info, pname)
			switch v := param.Value.(type) {
			case int:
				intParams[key] = flag.Int(key, v, param.Usage)
			case bool:
				boolParams[key] = flag.Bool(key, v, param.Usage)
			case string:
				stringParams[key] = flag.String(key, v, param.Usage)
			default:
				panic("unreachable") // Checked in AddChecker
			}
		}
	}

	p.checkerParams.ints = intParams
	p.checkerParams.bools = boolParams
	p.checkerParams.strings = stringParams

	return nil
}

func (p *program) checkerParamKey(info *linter.CheckerInfo, pname string) string {
	return "@" + info.Name + "." + pname
}

// bindDefaultEnabledList calculates the default value for -enable param.
func (p *program) bindDefaultEnabledList() error {
	var enabled []string
	for _, info := range p.infoList {
		enable := !info.HasTag("experimental") &&
			!info.HasTag("opinionated") &&
			!info.HasTag("performance")
		if enable {
			enabled = append(enabled, info.Name)
		}
	}
	p.filters.defaultCheckers = enabled
	return nil
}

func (p *program) parseArgs() error {
	flag.BoolVar(&p.filters.enableAll, "enableAll", false,
		`identical to -enable with all checkers listed. If true, -enable is ignored`)
	enable := flag.String("enable", strings.Join(p.filters.defaultCheckers, ","),
		`comma-separated list of enabled checkers. Can include #tags`)
	disable := flag.String("disable", "",
		`comma-separated list of checkers to be disabled. Can include #tags`)
	flag.IntVar(&p.exitCode, "exitCode", 1,
		`exit code to be used when lint issues are found`)
	flag.BoolVar(&p.checkTests, "checkTests", true,
		`whether to check test files`)
	flag.BoolVar(&p.shorterErrLocation, `shorterErrLocation`, true,
		`whether to replace error location prefix with $GOROOT and $GOPATH`)
	flag.BoolVar(&p.coloredOutput, `coloredOutput`, false,
		`whether to use colored output`)
	flag.BoolVar(&p.verbose, "v", false,
		`whether to print output useful during linter debugging`)

	flag.Parse()

	p.packages = flag.Args()
	p.filters.enable = strings.Split(*enable, ",")
	p.filters.disable = strings.Split(*disable, ",")

	if p.shorterErrLocation {
		wd, err := os.Getwd()
		if err != nil {
			log.Printf("getwd: %v", err)
		}
		p.workDir = addTrailingSlash(wd)
		p.gopath = addTrailingSlash(build.Default.GOPATH)
		p.goroot = addTrailingSlash(build.Default.GOROOT)
	}

	return nil
}

func addTrailingSlash(s string) string {
	if strings.HasSuffix(s, string(os.PathSeparator)) {
		return s
	}
	return s + string(os.PathSeparator)
}

// assignCheckerParams initializes checker parameter values using
// values that are coming from the command-line arguments.
func (p *program) assignCheckerParams() error {
	intParams := p.checkerParams.ints
	boolParams := p.checkerParams.bools
	stringParams := p.checkerParams.strings

	for _, info := range p.infoList {
		for pname, param := range info.Params {
			key := p.checkerParamKey(info, pname)
			switch param.Value.(type) {
			case int:
				info.Params[pname].Value = *intParams[key]
			case bool:
				info.Params[pname].Value = *boolParams[key]
			case string:
				info.Params[pname].Value = *stringParams[key]
			default:
				panic("unreachable") // Checked in AddChecker
			}
		}
	}

	return nil
}

var generatedFileCommentRE = regexp.MustCompile("Code generated .* DO NOT EDIT.")

func (p *program) isGenerated(f *ast.File) bool {
	return len(f.Comments) != 0 &&
		generatedFileCommentRE.MatchString(f.Comments[0].Text())
}

func (p *program) getFilename(f *ast.File) string {
	// See https://github.com/golang/go/issues/24498.
	return filepath.Base(p.fset.Position(f.Pos()).Filename)
}

func (p *program) shortenLocation(loc string) string {
	// If possible, construct relative path.
	relLoc := loc
	if p.workDir != "" {
		relLoc = strings.Replace(loc, p.workDir, "./", 1)
	}

	switch {
	case strings.HasPrefix(loc, p.gopath):
		loc = strings.Replace(loc, p.gopath, "$GOPATH"+string(os.PathSeparator), 1)
	case strings.HasPrefix(loc, p.goroot):
		loc = strings.Replace(loc, p.goroot, "$GOROOT"+string(os.PathSeparator), 1)
	}

	// Return the representation that is shorter.
	if len(relLoc) < len(loc) {
		return relLoc
	}
	return loc
}

func printWarning(p *program, rule, loc, warn string) {
	switch {
	case p.coloredOutput:
		log.Printf("%v: %v: %v\n",
			aurora.Magenta(aurora.Bold(loc)),
			aurora.Red(rule),
			warn)

	default:
		log.Printf("%s: %s: %s\n", loc, rule, warn)
	}
}

func loadPackages(cfg *packages.Config, patterns []string) ([]*packages.Package, error) {
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, err
	}

	result := pkgs[:0]
	pkgload.VisitUnits(pkgs, func(u *pkgload.Unit) {
		if u.ExternalTest != nil {
			result = append(result, u.ExternalTest)
		}

		if u.Test != nil {
			// Prefer tests to the base package, if present.
			result = append(result, u.Test)
		} else {
			result = append(result, u.Base)
		}
	})
	return result, nil
}
