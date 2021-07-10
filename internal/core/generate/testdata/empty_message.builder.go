package testdata

// TeaBuilder provides a fluent api for building an instance of Tea
type TeaBuilder struct {
    tea *Tea
}

// NewTea starts building the Tea
func NewTea() *TeaBuilder {
    return &TeaBuilder{&Tea{}}
}

// Build builds the Tea
func (teaBuilder *TeaBuilder) Build() *Tea {
    return teaBuilder.tea
}