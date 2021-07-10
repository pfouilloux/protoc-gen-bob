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
	{{- $lower_name := toLower .Name -}}
    {{- $arg := (printf "%s %s" $lower_name .Kind) -}}
    {{- $val := $lower_name -}}
    {{- if and (eq .Kind "bool") (not .IsOptional) -}}
        {{- $arg = "" -}}
        {{- $val = "true" -}}
    {{- else if .IsMessage -}}
		{{- $arg = (printf "%s *%sBuilder" $lower_name .Kind) -}}
		{{- $val = (printf "%s.Build()" $lower_name) -}}
	{{- else if .IsOptional -}}
		{{- $val = (printf "&%s" $lower_name)}}
    {{- end}}
// {{.Name}} sets {{$name}}.{{.Name}}
func ({{$receiver}} *{{$builder}}) {{.Name}}({{$arg}}) *{{$builder}} {
    {{$receiver}}.{{$backing}}.{{.Name}} = {{$val}}
    return {{$receiver}}
}
{{end}}
// Build builds the {{.Name}}
func ({{$receiver}} *{{$builder}}) Build() *{{.Name}} {
    return {{$receiver}}.{{$backing}}
}

{{- end}}