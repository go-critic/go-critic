package analyzer

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/go-critic/go-critic/checkers" // Register go-critic checkers
	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/analysis"
)

type gocritic struct {
	infoList  []*linter.CheckerInfo
	goVersion linter.GoVersion
}

var (
	globalGocriticMu        sync.Mutex
	globalGocritic          *gocritic
	globalInitErrorReported bool
)

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	critic, err := prepareGocritic()
	if err != nil {
		return nil, fmt.Errorf("init error: %w", err)
	}

	ctx := linter.NewContext(pass.Fset, pass.TypesSizes)
	ctx.GoVersion = critic.goVersion
	ctx.SetPackageInfo(pass.TypesInfo, pass.Pkg)

	checkers, err := critic.createCheckers(ctx)
	if err != nil {
		return nil, err
	}

	for _, f := range pass.Files {
		f := f
		filename := filepath.Base(pass.Fset.Position(f.Pos()).Filename)
		ctx.SetFileInfo(filename, f)

		for _, c := range checkers {
			warnings := c.Check(f)
			for _, warning := range warnings {
				pass.Report(asDiag(c, warning))
			}
		}
	}

	return nil, nil
}

func asDiag(c *linter.Checker, warning linter.Warning) analysis.Diagnostic {
	diag := analysis.Diagnostic{
		Pos:     warning.Pos,
		Message: fmt.Sprintf("%s: %s", c.Info.Name, warning.Text),
	}

	if warning.HasQuickFix() {
		diag.SuggestedFixes = []analysis.SuggestedFix{
			{
				Message: "suggested replacement",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     warning.Suggestion.From,
						End:     warning.Suggestion.To,
						NewText: warning.Suggestion.Replacement,
					},
				},
			},
		}
	}
	return diag
}

// prepareGocritic initializes a new gocritic object,
// but unlike newGocritic() it could use a cached version.
func prepareGocritic() (*gocritic, error) {
	if DisableCache {
		return newGocritic()
	}

	globalGocriticMu.Lock()
	defer globalGocriticMu.Unlock()

	// Don't report init error ever again if it was already reported.
	if globalInitErrorReported {
		return nil, nil
	}

	if globalGocritic != nil {
		return globalGocritic, nil
	}

	critic, err := newGocritic()
	if err != nil {
		globalInitErrorReported = true
		return nil, err
	}
	globalGocritic = critic
	return critic, nil
}

func newGocritic() (*gocritic, error) {
	critic := &gocritic{
		infoList: filterCheckersList(registeredCheckers),
	}

	ver, err := linter.ParseGoVersion(flagGoVersion)
	if err != nil {
		return nil, err
	}
	critic.goVersion = ver

	for _, info := range critic.infoList {
		for pname, param := range info.Params {
			key := checkerParamName(info, pname)
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

	return critic, nil
}

func filterCheckersList(infoList []*linter.CheckerInfo) []*linter.CheckerInfo {
	parseKeys := func(keys []string, byName, byTag map[string]bool) {
		for _, key := range keys {
			if strings.HasPrefix(key, "#") {
				byTag[key[len("#"):]] = true
			} else {
				byName[key] = true
			}
		}
	}
	splitValues := func(s string) []string {
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}

	disableArg := flagDisable
	if disableArg == "<default>" {
		if flagEnableAll {
			disableArg = ""
		} else {
			disableArg = "#experimental,#opinionated,#performance"
		}
	}

	enabledByName := make(map[string]bool)
	enabledTags := make(map[string]bool)
	parseKeys(splitValues(flagEnable), enabledByName, enabledTags)
	disabledByName := make(map[string]bool)
	disabledTags := make(map[string]bool)
	parseKeys(splitValues(disableArg), disabledByName, disabledTags)

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

	var filtered []*linter.CheckerInfo

	for _, info := range infoList {
		enabled := flagEnableAll || enabledByName[info.Name] || enabledByTag(info)
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

		if flagDebugInit && !enabled {
			log.Printf("\tdebug: %s: %s", info.Name, notice)
		}
		if !enabled {
			continue
		}
		filtered = append(filtered, info)
	}
	if flagDebugInit {
		for _, info := range filtered {
			log.Printf("\tdebug: %s is enabled", info.Name)
		}
	}

	return filtered
}

func (critic *gocritic) createCheckers(ctx *linter.Context) ([]*linter.Checker, error) {
	checkers := make([]*linter.Checker, len(critic.infoList))
	for i, info := range critic.infoList {
		c, err := linter.NewChecker(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("init %s: %w", info.Name, err)
		}
		checkers[i] = c
	}
	return checkers, nil
}
