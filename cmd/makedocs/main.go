package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-critic/go-critic/checkers"
	"github.com/go-critic/go-critic/linter"
)

const (
	docsPath      = "../../docs/"
	templatesPath = docsPath + "templates/"
)

func main() {
	if err := checkers.InitEmbeddedRules(); err != nil {
		panic(err)
	}

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
	
	// A bit hacky but works, see https://github.com/go-critic/go-critic/pull/1403.
	buff := bytes.ReplaceAll(buf.Bytes(), []byte("<all>"), []byte("&lt;all&gt;"))

	if err := os.WriteFile(docsPath+"overview.md", buff, 0o600); err != nil {
		log.Fatalf("write output file: %v", err)
	}
}

func parseTemplate(names ...string) *template.Template {
	paths := make([]string, len(names))
	for i := range names {
		paths[i] = templatesPath + names[i]
	}
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}
	return template.Must(template.New("overview").Funcs(funcMap).ParseFiles(paths...))
}
