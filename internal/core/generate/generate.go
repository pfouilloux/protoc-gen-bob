package generate

import (
	"embed"
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"io"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFs embed.FS

var templates = template.Must(template.New("builder.go.tpl").
	Funcs(template.FuncMap{"toLower": strings.ToLower}).
	ParseFS(templateFs, "templates/*.go.tpl"))

type Spec struct {
	name   string
	input  model.File
	writer io.Writer
}

func NewSpec(name string, file model.File, wr io.Writer) Spec { return Spec{name, file, wr} }
func (s Spec) Writer() io.Writer                              { return s.writer }

func Builders(specs ...Spec) error {
	// TODO: Run these in parallel?
	for i, task := range specs {
		if err := templates.Execute(task.writer, task.input); err != nil {
			return fmt.Errorf("failed create a builder for task %d '%s': %v", i, task.name, err)
		}
	}

	return nil
}
