package builder

type ignooBuilder struct {
	windowType string
	doorType string
	floorType int
}

func newignooBuilder() *ignooBuilder {
	return &ignooBuilder{}
}

func (b *ignooBuilder) setwindowType() {
	b.windowType = "snow window "
}
func (b *ignooBuilder) setdoorType() {
	b.doorType = "glass door"
}
func (b *ignooBuilder) setfloorType() {
	b.floorType = 1
}
func (b *ignooBuilder) getHouse() House {
	return House { 
		windowType: b.windowType, 
		doorType: b.doorType,
		floorType: b.floorType,
	}
}
