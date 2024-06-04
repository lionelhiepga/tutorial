package builder

/*director is a struct*/
type Director struct {
	builder IBuilder
}

/*newdirector is a function*/
func NewDirector(b IBuilder) *Director{
	return &Director{
		builder: b,
	}
}

/*setbuilder is a function*/
func (d *Director) SetBuilder(b IBuilder){
	d.builder = b
} 

/*builderHouse is a function*/ //hàm khởi tạo gọi từ interface
func (d *Director) SetBuilderHouse() House {
	d.builder.setdoorType() 
	d.builder.setwindowType()
	d.builder.setfloorType()
	return d.builder.getHouse()
}