package builder

type normalBuilder struct {
	windowType string
	doorType string
	floorType int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setwindowType() {
	b.windowType = "wooden window "
}
func (b *normalBuilder) setdoorType() {
	b.doorType = "wood door "
}
func (b *normalBuilder) setfloorType() {
	b.floorType = 2
}
func (b *normalBuilder) getHouse() House { 
	return House { 
		windowType: b.windowType, 
		doorType: b.doorType,
		floorType: b.floorType,
	}
}