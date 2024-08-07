//go:build generate

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/quasilyte/go-ruleguard/ruleguard/irconv"
	"github.com/quasilyte/go-ruleguard/ruleguard/irprint"
)

// This program generates a loadable IR for ruleguard
// so we don't have to load the rules from AST and typecheck
// them every time.

func main() {
	log.SetFlags(0)
	if err := precompile(); err != nil {
		log.Printf("error: %v", err)
	}
}

func precompile() error {
	flagRules := flag.String("rules", "", "path to a ruleguard rules file")
	flagOutput := flag.String("o", "", "output file name")
	flag.Parse()

	fset := token.NewFileSet()
	filename := strings.TrimSpace(*flagRules)
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read %s: %v", filename, err)
	}
	r := bytes.NewReader(fileData)
	parserFlags := parser.ParseComments
	f, err := parser.ParseFile(fset, filename, r, parserFlags)
	if err != nil {
		return fmt.Errorf("parse %s: %v", filename, err)
	}
	imp := importer.For("source", nil)
	typechecker := types.Config{Importer: imp}
	types := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Uses:  map[*ast.Ident]types.Object{},
		Defs:  map[*ast.Ident]types.Object{},
	}
	pkg, err := typechecker.Check("gorules", fset, []*ast.File{f}, types)
	if err != nil {
		return fmt.Errorf("typecheck %s: %v", filename, err)
	}
	irconvCtx := &irconv.Context{
		Pkg:   pkg,
		Types: types,
		Fset:  fset,
		Src:   fileData,
	}
	irfile, err := irconv.ConvertFile(irconvCtx, f)
	if err != nil {
		return fmt.Errorf("compile %s: %v", filename, err)
	}

	var rulesText bytes.Buffer
	irprint.File(&rulesText, irfile)

	fileTemplate := template.Must(template.New("gorules").Parse(`// Code generated by "precompile.go". DO NOT EDIT.

package rulesdata

import "github.com/quasilyte/go-ruleguard/ruleguard/ir"

var PrecompiledRules = &{{$.RulesText}}
`))

	var generated bytes.Buffer
	err = fileTemplate.Execute(&generated, map[string]interface{}{
		"RulesText": rulesText.String(),
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(*flagOutput, generated.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
