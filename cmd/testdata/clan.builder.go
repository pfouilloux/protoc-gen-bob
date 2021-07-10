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

// WargBuilder provides a fluent api for building an instance of Warg
type WargBuilder struct {
	warg *Warg
}

// NewWarg starts building the Warg
func NewWarg() *WargBuilder {
	return &WargBuilder{&Warg{}}
}

// Name sets Warg.Name
func (wargBuilder *WargBuilder) Name(name string) *WargBuilder {
	wargBuilder.warg.Name = name
	return wargBuilder
}

// Breed sets Warg.Breed
func (wargBuilder *WargBuilder) Breed(breed Warg_Breed) *WargBuilder {
	wargBuilder.warg.Breed = breed
	return wargBuilder
}

// Saddled sets Warg.Saddled
func (wargBuilder *WargBuilder) Saddled() *WargBuilder {
	wargBuilder.warg.Saddled = true
	return wargBuilder
}

// Build builds the Warg
func (wargBuilder *WargBuilder) Build() *Warg {
	return wargBuilder.warg
}
