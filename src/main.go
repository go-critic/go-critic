package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

type logger interface {
	log(format string, args ...interface{}) error
}

type checker interface {
	run(fset *token.FileSet, logger logger) error
}

type consoleLogger struct{}

func (c consoleLogger) log(format string, args ...interface{}) error {
	_, err := fmt.Printf(format, args...)
	return err
}

var checkers = []checker{}

func main() {
	fset := token.NewFileSet()
	dir := flag.String("dir", "", "project directory")
	flag.Parse()

	if *dir == "" {
		flag.Usage()
		os.Exit(1)
	}

	parser.ParseDir(fset, *dir, nil, parser.ParseComments)

	cl := consoleLogger{}

	for _, c := range checkers {
		c.run(fset, &cl)
	}
}
