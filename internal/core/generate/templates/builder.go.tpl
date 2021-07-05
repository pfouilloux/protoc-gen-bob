package {{.Pkg}}

{{range .Messages -}}
{{$name := .Name -}}
{{$backing := toLower .Name -}}
{{$receiver := print $backing "Builder" -}}
{{$builder := print .Name "Builder" -}}

// {{$builder}} provides a fluent api for building an instance of {{.Name}}
type {{$builder}} struct {
    {{$backing}} *{{.Name}}
}

// New{{.Name}} starts building the {{.Name}}
func New{{.Name}}() *{{$builder}} {
    return &{{$builder}}{&{{.Name}}{}}
}
{{range .Fields}}
// {{.Name}} sets {{$name}}.{{.Name}}
func ({{$receiver}} *{{$builder}}) {{.Name}}({{toLower .Name}} {{.Kind}}) *{{$builder}} {
    {{$receiver}}.{{$backing}}.{{.Name}} = {{if .IsOptional}}&{{end}}{{toLower .Name}}
    return {{$receiver}}
}
{{end}}
// Build builds the {{.Name}}
func ({{$receiver}} *{{$builder}}) Build() *{{.Name}} {
    return {{$receiver}}.{{$backing}}
}

{{- end}}