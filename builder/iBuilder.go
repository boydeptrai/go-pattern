package builder


/*IBuilder is interface*/
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

/*GetBuilder is function*/
func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
		
	}
	return nil
}