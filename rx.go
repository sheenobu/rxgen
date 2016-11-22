// Package rx is the reactive typewriter
package rx

import (
	"io"
	"strings"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(&Typewriter{})
	if err != nil {
		panic(err)
	}
}

// Typewriter is the typewriter for reactive types
type Typewriter struct {
}

// Name returns the name of the typewriter
func (rx *Typewriter) Name() string {
	return "rx"
}

// Imports returns the required imports for the generated file
func (rx *Typewriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	result = append(result, typewriter.ImportSpec{Name: "sync", Path: "sync"})
	return
}

// Data is the template data
type Data struct {
	Name string
	Type string
}

// Write writes the file contents
func (rx *Typewriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(rx)

	if !found {
		// nothing to be done
		return nil
	}

	typeName := t.Name
	prefix := "Rx"

	for _, val := range tag.Values {
		if val.Name == "Builtin" {
			prefix = ""
			typeName = val.TypeParameters[0].Name
			t.Name = strings.Trim(t.Name, "_")
		}
	}

	data := Data{
		Name: prefix + t.Name,
		Type: typeName,
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}
