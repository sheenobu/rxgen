package main

import (
	"os"
	"strings"

	"text/template"
)

type rxData struct {
	Invocation        string
	AdditionalImports string

	Package string
	Type    string
	Name    string
}

func (g *Generator) generateRx(name string, objectType string) {

	data := rxData{
		Invocation: strings.Join(os.Args[1:], " "),
		Package:    g.pkg.name,
		Type:       objectType,
		Name:       name,
	}

	t := template.Must(template.New("rx").Parse(rxtemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
