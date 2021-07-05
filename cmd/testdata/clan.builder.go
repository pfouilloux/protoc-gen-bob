package testdata

// OrcBuilder provides a fluent api for building an instance of Orc
type OrcBuilder struct {
	orc *Orc
}

// NewOrc starts building the Orc
func NewOrc() *OrcBuilder {
	return &OrcBuilder{&Orc{}}
}

// Name sets Orc.Name
func (orcBuilder *OrcBuilder) Name(name string) *OrcBuilder {
	orcBuilder.orc.Name = name
	return orcBuilder
}

// Age sets Orc.Age
func (orcBuilder *OrcBuilder) Age(age uint32) *OrcBuilder {
	orcBuilder.orc.Age = age
	return orcBuilder
}

// Title sets Orc.Title
func (orcBuilder *OrcBuilder) Title(title string) *OrcBuilder {
	orcBuilder.orc.Title = &title
	return orcBuilder
}

// Build builds the Orc
func (orcBuilder *OrcBuilder) Build() *Orc {
	return orcBuilder.orc
}
