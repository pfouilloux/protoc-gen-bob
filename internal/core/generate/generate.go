package generate

import (
	"embed"
	_ "embed"
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

type Task struct {
	name   string
	input  model.File
	writer io.Writer
}

func NewTask(name string, file model.File, wr io.Writer) Task { return Task{name, file, wr} }

func Builders(tasks ...Task) error {
	for i, task := range tasks {
		if err := templates.Execute(task.writer, task.input); err != nil {
			return fmt.Errorf("failed create a builder for task %d '%s': %v", i, task.name, err)
		}
	}

	return nil
}
