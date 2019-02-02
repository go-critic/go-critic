package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	_ "github.com/go-critic/go-critic/checkers"
	"github.com/go-lintpack/lintpack"
)

const (
	docsPath      = "../../docs/"
	templatesPath = docsPath + "templates/"
)

func main() {
	tmpl := template.Must(template.ParseFiles(templatesPath + "overview.md.tmpl"))
	buf := bytes.Buffer{}
	err := tmpl.Execute(&buf, struct {
		Checkers []*lintpack.CheckerInfo
	}{
		Checkers: lintpack.GetCheckersInfo(),
	})
	if err != nil {
		log.Fatalf("render template: %v", err)
	}
	if err := ioutil.WriteFile(docsPath+"overview.md", buf.Bytes(), 0600); err != nil {
		log.Fatalf("write output file: %v", err)
	}
}
