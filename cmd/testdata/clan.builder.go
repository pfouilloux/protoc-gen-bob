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

// CaragorBuilder provides a fluent api for building an instance of Caragor
type CaragorBuilder struct {
	caragor *Caragor
}

// NewCaragor starts building the Caragor
func NewCaragor() *CaragorBuilder {
	return &CaragorBuilder{&Caragor{}}
}

// Name sets Caragor.Name
func (caragorBuilder *CaragorBuilder) Name(name string) *CaragorBuilder {
	caragorBuilder.caragor.Name = name
	return caragorBuilder
}

// Equipment sets Caragor.Equipment
func (caragorBuilder *CaragorBuilder) Equipment(equipment *Caragor_EquipmentBuilder) *CaragorBuilder {
	caragorBuilder.caragor.Equipment = equipment.Build()
	return caragorBuilder
}

// Armoured sets Caragor.Armoured
func (caragorBuilder *CaragorBuilder) Armoured() *CaragorBuilder {
	caragorBuilder.caragor.Armoured = true
	return caragorBuilder
}

// Build builds the Caragor
func (caragorBuilder *CaragorBuilder) Build() *Caragor {
	return caragorBuilder.caragor
}

// Caragor_EquipmentBuilder provides a fluent api for building an instance of Caragor_Equipment
type Caragor_EquipmentBuilder struct {
	caragor_equipment *Caragor_Equipment
}

// NewCaragor_Equipment starts building the Caragor_Equipment
func NewCaragor_Equipment() *Caragor_EquipmentBuilder {
	return &Caragor_EquipmentBuilder{&Caragor_Equipment{}}
}

// Spears sets Caragor_Equipment.Spears
func (caragor_equipmentBuilder *Caragor_EquipmentBuilder) Spears(spears int32) *Caragor_EquipmentBuilder {
	caragor_equipmentBuilder.caragor_equipment.Spears = spears
	return caragor_equipmentBuilder
}

// Bows sets Caragor_Equipment.Bows
func (caragor_equipmentBuilder *Caragor_EquipmentBuilder) Bows(bows int32) *Caragor_EquipmentBuilder {
	caragor_equipmentBuilder.caragor_equipment.Bows = bows
	return caragor_equipmentBuilder
}

// Provisions sets Caragor_Equipment.Provisions
func (caragor_equipmentBuilder *Caragor_EquipmentBuilder) Provisions(provisions int32) *Caragor_EquipmentBuilder {
	caragor_equipmentBuilder.caragor_equipment.Provisions = provisions
	return caragor_equipmentBuilder
}

// Build builds the Caragor_Equipment
func (caragor_equipmentBuilder *Caragor_EquipmentBuilder) Build() *Caragor_Equipment {
	return caragor_equipmentBuilder.caragor_equipment
}
