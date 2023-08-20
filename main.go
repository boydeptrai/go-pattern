package main

import (
	"fmt"

	p "github.com/boydeptrai/go-pattern/prototype"
	"github.com/boydeptrai/go-pattern/abstractfactory"
	"github.com/boydeptrai/go-pattern/builder"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildeHouse()

	//Abstract Factory
	adidasFactory := abstractfactory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	printShoeDetails(adidasShoe)
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(adidasShort)

	nikeFactory := abstractfactory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	printShoeDetails(nikeShoe)
	nikeShort := nikeFactory.MakeShort()
	printShortDetails(nikeShort)
    

	// Prototype
	file1 := &p.File{Name: "File 1"}
	file2 := &p.File{Name: "File 2"}
	file3 := &p.File{Name: "File 3"}
	folder1 := &p.Folder{
		Childrens: []p.INode{file1},
		Name: "Folder 1",
	}
	folder2 := &p.Folder {
		Childrens: []p.INode{folder1, file2, file3},
		Name: "Folder 2",
	}

	fmt.Println("\n Printing for Folder 2")
	folder2.Print(" ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting for close folder 2")
	cloneFolder.Print("  ")


	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type:%s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetNumFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildeHouse()
    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetNumFloor())
}

func printShoeDetails(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetails(s abstractfactory.IShort) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
