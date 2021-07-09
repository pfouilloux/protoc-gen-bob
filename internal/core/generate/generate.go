package generate

import (
	"embed"
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"io"
	"strings"
	"text/template"
)

type Generator func(wr io.Writer, file model.File) error

func Generate(wr io.Writer, file model.File) error {
	if err := templates.Execute(wr, file); err != nil {
		return fmt.Errorf("failed create a builder: %v", err)
	}

	return nil
}

//go:embed templates/*
var templateFs embed.FS

var templates = template.Must(template.New("builder.go.tpl").
	Funcs(template.FuncMap{"toLower": strings.ToLower}).
	ParseFS(templateFs, "templates/*.go.tpl"))
