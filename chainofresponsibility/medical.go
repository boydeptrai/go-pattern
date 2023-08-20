package chainofresponsibility

import (
	"fmt"
)

/*Medical is struct */
type Medical struct {
	next Department
}

/*Execute is fuction*/
func (m *Medical) Execute(p *Patient) {
	if p.isMedicineProvied {
		fmt.Println("Patient already provided medicine")
		m.next.Execute(p)
		return
	}
	fmt.Println("We are providing medicine to patient")
	p.isMedicineProvied = true
	m.next.Execute(p)
}
/*SetNext is fuction*/
func (m *Medical) SetNext(next Department) {
	m.next = next
}