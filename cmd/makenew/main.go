package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const (
	checkersPath = "../../lint/"
)

func main() {
	name := flag.String("name", "", "name of a new checker")
	flag.Parse()

	if *name == "" {
		log.Fatalf("empty checker name")
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
	if err := os.MkdirAll(testDir, os.ModePerm); err != nil {
		log.Fatalf("cannot create a test dir: %v", err)
	}
	if _, err := os.OpenFile(testDir+"/positive_tests.go", os.O_RDONLY|os.O_CREATE, os.ModePerm); err != nil {
		log.Fatalf("cannot create a test file: %v", err)
	}
	if _, err := os.OpenFile(testDir+"/negative_tests.go", os.O_RDONLY|os.O_CREATE, os.ModePerm); err != nil {
		log.Fatalf("cannot create a test file: %v", err)
	}

}
