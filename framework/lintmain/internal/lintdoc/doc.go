package lintdoc

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/go-critic/go-critic/framework/linter"
)

// Main implements sub-command entry point.
func Main() {
	flag.Parse()

	switch args := flag.Args(); len(args) {
	case 0:
		printShortDoc()
	case 1:
		printDoc(args[0])
	default:
		log.Fatalf("expected 0 or 1 positional arguments")
	}
}

func printShortDoc() {
	for _, info := range linter.GetCheckersInfo() {
		fmt.Printf("%s %v\n", info.Name, info.Tags)
	}
}

func printDoc(name string) {
	info := findInfoByName(name)
	if info == nil {
		log.Fatalf("checker with name %q not found", name)
	}

	tmplString := `{{.Checker.Name}} checker documentation
URL: {{.Checker.Collection.URL}}
Tags: {{.Checker.Tags}}

{{.Checker.Summary}}.
{{ if .Checker.Details }}
{{.Checker.Details}}
{{ end }}
Non-compliant code:
{{.Checker.Before}}

Compliant code:
{{.Checker.After}}
{{- if .Checker.Note }}

{{.Checker.Note}}
{{- end }}
{{- if .Checker.Params }}

Checker parameters:
{{- range $key, $_ := .Checker.Params }}
  -@{{$.Checker.Name}}.{{$key}} {{index $.ParamTypes $key}}
    	{{.Usage}} (default {{.Value}})
{{- end }}
{{- end }}
`

	var templateData struct {
		Checker    *linter.CheckerInfo
		ParamTypes map[string]string
	}
	templateData.Checker = info
	templateData.ParamTypes = make(map[string]string)
	for pname, p := range info.Params {
		templateData.ParamTypes[pname] = fmt.Sprintf("%T", p.Value)
	}

	tmpl := template.Must(template.New("doc").Parse(tmplString))
	if err := tmpl.Execute(os.Stdout, templateData); err != nil {
		panic(fmt.Sprintf("executing checker doc template: %v", err))
	}
}

func findInfoByName(name string) *linter.CheckerInfo {
	for _, info := range linter.GetCheckersInfo() {
		if info.Name == name {
			return info
		}
	}
	return nil
}
