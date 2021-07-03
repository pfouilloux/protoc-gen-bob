package testdata

// LocationBuilder provides a fluent api for building testdata.Location
type LocationBuilder struct {
	location *Location
}

// NewLocation starts building the testdata.Location
func NewLocation(long int32, lat int32) *LocationBuilder {
	return &LocationBuilder{&Location{Long: long, Lat: lat}}
}

// Build builds the testdata.Location
func (x *LocationBuilder) Build() *Location {
	return x.location
}
