package testdata

// KettleBuilder provides a fluent api for building an instance of Kettle
type KettleBuilder struct {
    kettle *Kettle
}

// NewKettle starts building the Kettle
func NewKettle() *KettleBuilder {
    return &KettleBuilder{&Kettle{}}
}

// Material sets Kettle.Material
func (kettleBuilder *KettleBuilder) Material(material string) *KettleBuilder {
    kettleBuilder.kettle.Material = material
    return kettleBuilder
}

// Colour sets Kettle.Colour
func (kettleBuilder *KettleBuilder) Colour(colour string) *KettleBuilder {
    kettleBuilder.kettle.Colour = &colour
    return kettleBuilder
}

// Capacity sets Kettle.Capacity
func (kettleBuilder *KettleBuilder) Capacity(capacity uint32) *KettleBuilder {
    kettleBuilder.kettle.Capacity = capacity
    return kettleBuilder
}

// Build builds the Kettle
func (kettleBuilder *KettleBuilder) Build() *Kettle {
    return kettleBuilder.kettle
}