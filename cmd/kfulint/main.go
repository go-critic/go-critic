package main

import (
	"flag"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/PieselBois/kfulint/lint"
)

func checkers() []lint.Checker {
	return []lint.Checker{}
}

func blame(format string, args ...interface{}) {
	log.Printf(format, args...)
	flag.Usage()
	os.Exit(1)
}

// parseArgv processes command-line arguments and fills ctx argument with them.
// Terminates program on error.
func parseArgv(ctx *lint.Context) {
	flag.StringVar(&ctx.PkgDir, "dir", "", "package directory")

	if ctx.PkgDir == "" {
		blame("Illegal empty -dir argument\n")
	}

	flag.Parse()
}

func parsePackage(ctx *lint.Context) {
	// TODO: save ParseDir results into ctx.
	parser.ParseDir(ctx.Fset, ctx.PkgDir, nil, parser.ParseComments)
}

func main() {
	ctx := lint.Context{
		Fset: token.NewFileSet(),
	}

	parseArgv(&ctx)
	parsePackage(&ctx)

	for _, c := range checkers() {
		if err := c.Run(&ctx); err != nil {
			log.Print(err)
		}
	}
}
