package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const (
	checkersPath = "../../lint/"
)

func main() {
	name := flag.String("name", "", "")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please, provide a name for a new checker")
		os.Exit(2)
	}

	tmpl := template.Must(template.ParseFiles("new_checker.go.tmpl"))

	buf := bytes.Buffer{}
	err := tmpl.Execute(&buf, struct {
		Name string
	}{
		Name: *name,
	})
	if err != nil {
		log.Fatalf("render template: %v", err)
	}

	filename := checkersPath + *name + "_checker.go"

	if err := ioutil.WriteFile(filename, buf.Bytes(), 0600); err != nil {
		log.Fatalf("write output file: %v", err)
	}

	testDir := checkersPath + "testdata/" + *name
	os.MkdirAll(testDir, os.ModePerm)
	os.OpenFile(testDir+"/positive_tests.go", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	os.OpenFile(testDir+"/negative_tests.go", os.O_RDONLY|os.O_CREATE, os.ModePerm)

}
