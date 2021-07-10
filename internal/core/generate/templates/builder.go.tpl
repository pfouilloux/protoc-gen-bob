package {{.Pkg -}}

{{range .Messages -}}
{{$name := .Name -}}
{{$backing := toLower .Name -}}
{{$receiver := print $backing "Builder" -}}
{{$builder := print .Name "Builder"}}

// {{$builder}} provides a fluent api for building an instance of {{.Name}}
type {{$builder}} struct {
    {{$backing}} *{{.Name}}
}

// New{{.Name}} starts building the {{.Name}}
func New{{.Name}}() *{{$builder}} {
    return &{{$builder}}{&{{.Name}}{}}
}
{{range .Fields -}}
{{$isReqBool := and (eq .Kind "bool") (ne .IsOptional true)}}
// {{.Name}} sets {{$name}}.{{.Name}}
func ({{$receiver}} *{{$builder}}) {{.Name}}({{if not $isReqBool}}{{toLower .Name}} {{.Kind}}{{end}}) *{{$builder}} {
    {{$receiver}}.{{$backing}}.{{.Name}} = {{if $isReqBool}}true{{else}}{{if .IsOptional}}&{{end}}{{toLower .Name}}{{end}}
    return {{$receiver}}
}
{{end}}
// Build builds the {{.Name}}
func ({{$receiver}} *{{$builder}}) Build() *{{.Name}} {
    return {{$receiver}}.{{$backing}}
}

{{- end}}