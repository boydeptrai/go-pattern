package builder

/*Director is struct*/
type Director struct {
	builder IBuilder
}

/*NewDirector is function*/
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

/*SetBuilder is function*/
func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

/*BuildeHouse is function*/
func (d *Director) BuildeHouse() House {
	d.builder.setDoorType()
	d.builder.setNumFloor()
	d.builder.setWindowType()
	return d.builder.getHouse()
}