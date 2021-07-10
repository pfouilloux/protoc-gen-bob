package testdata

// CupBuilder provides a fluent api for building an instance of Cup
type CupBuilder struct {
    cup *Cup
}

// NewCup starts building the Cup
func NewCup() *CupBuilder {
    return &CupBuilder{&Cup{}}
}

// Tea sets Cup.Tea
func (cupBuilder *CupBuilder) Tea(tea *TeaBuilder) *CupBuilder {
    cupBuilder.cup.Tea = tea.Build()
    return cupBuilder
}

// Sugar sets Cup.Sugar
func (cupBuilder *CupBuilder) Sugar(sugar *SugarBuilder) *CupBuilder {
    cupBuilder.cup.Sugar = sugar.Build()
    return cupBuilder
}

// Build builds the Cup
func (cupBuilder *CupBuilder) Build() *Cup {
    return cupBuilder.cup
}