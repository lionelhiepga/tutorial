package builder

/*iBuilder interface*/
type IBuilder interface {
	setwindowType()
	setdoorType()
	setfloorType()
	getHouse() House  //phương thức trả về 1 cái house
}

//getbuilder in func
func getBuilder(builderType string) IBuilder {
	switch builderType {
		case "normal": 
			return &normalBuilder{}
		case "igloo": 
			return &ignooBuilder{}
	}
	return nil
} 

