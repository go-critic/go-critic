package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	_ "github.com/go-critic/go-critic/checkers"
	"github.com/go-critic/go-critic/framework/linter"
)

const (
	docsPath      = "../../docs/"
	templatesPath = docsPath + "templates/"
)

func main() {
	tmpl := parseTemplate(
		"overview.md.tmpl",
		"checker.partial.tmpl",
		"checker_tr.partial.tmpl")

	buf := bytes.Buffer{}
	err := tmpl.ExecuteTemplate(&buf, "overview", struct {
		Checkers []*linter.CheckerInfo
	}{
		Checkers: linter.GetCheckersInfo(),
	})
	if err != nil {
		log.Fatalf("render template: %v", err)
	}
	if err := os.WriteFile(docsPath+"overview.md", buf.Bytes(), 0o600); err != nil {
		log.Fatalf("write output file: %v", err)
	}
}

func parseTemplate(names ...string) *template.Template {
	paths := make([]string, len(names))
	for i := range names {
		paths[i] = templatesPath + names[i]
	}
	return template.Must(template.ParseFiles(paths...))
}
