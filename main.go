package main

import (
	"github.com/boydeptrai/go-pattern/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildeHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type:%s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetNumFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildeHouse()
    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetNumFloor())
}
