package prototype

import (
	"fmt"
)

type Folder struct {
	Childrens []INode
	Name string
}
/* Print is function*/
func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, i := range f.Childrens {
		i.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name  + "_Clone"}
	var tempChildrens []INode
	for _, i := range f.Childrens {
		copy := i.Clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.Childrens = tempChildrens
	return cloneFolder
}