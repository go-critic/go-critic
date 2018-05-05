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

func main() {
	fset := token.NewFileSet()
	dir := flag.String("dir", "", "project directory")
	flag.Parse()

	if *dir == "" {
		blame("Illegal empty -dir argument\n")
	}

	parser.ParseDir(fset, *dir, nil, parser.ParseComments)

	ct := lint.Context{}
	ct.FileSet = fset
	ct.Flags = nil

	for _, c := range checkers() {

		if err := c.Run(&ct); err != nil {
			log.Print(err)
		}
	}
}
