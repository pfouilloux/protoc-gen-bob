package testdata

// CupBuilder provides a fluent api for building an instance of Cup
type CupBuilder struct {
    cup *Cup
}

// NewCup starts building the Cup
func NewCup() *CupBuilder {
    return &CupBuilder{&Cup{}}
}

// Full sets Cup.Full
func (cupBuilder *CupBuilder) Full() *CupBuilder {
    cupBuilder.cup.Full = true
    return cupBuilder
}

// Hot sets Cup.Hot
func (cupBuilder *CupBuilder) Hot(hot bool) *CupBuilder {
    cupBuilder.cup.Hot = &hot
    return cupBuilder
}

// Build builds the Cup
func (cupBuilder *CupBuilder) Build() *Cup {
    return cupBuilder.cup
}