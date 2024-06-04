package builder


/*house is struct */
type House struct{
	windowType string
	doorType string
	floorType int
}

/*getwindowType is getter, các hàm trả về có thể được gọi ở các package khác*/

func (h *House) GetwindowType() string {
	return h.windowType
}
func (h *House) GetdoorType() string {
	return h.doorType
}
func (h *House) Getfloor() int {
	return h.floorType
}

