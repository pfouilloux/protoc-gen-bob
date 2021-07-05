package model

type File struct {
	pkg      string
	messages []Message
}

func NewFile(pkg string, messages ...Message) File { return File{pkg, messages} }
func (f File) Pkg() string                         { return f.pkg }
func (f File) Messages() []Message                 { return f.messages }

type Message struct {
	name   string
	fields []Field
}

func NewMessage(name string, fields ...Field) Message { return Message{name, fields} }
func (m Message) Name() string                        { return m.name }
func (m Message) Fields() []Field                     { return m.fields }

type Field struct {
	name     string
	kind     string
	optional bool
}

func NewField(name, kind string, optional bool) Field { return Field{name, kind, optional} }
func (f Field) IsOptional() bool                      { return f.optional }
func (f Field) Name() string                          { return f.name }
func (f Field) Kind() string                          { return f.kind }
