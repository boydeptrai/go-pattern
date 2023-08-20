package chainofresponsibility

import (
	"fmt"
)

type Doctor struct {
	next Department
}
/*Execute is function*/
func (d *Doctor) Execute(p *Patient) {
	if p.isDoctorChecked {
		fmt.Println("Patient already checked by doctor")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor is checking patient")
	d.next.Execute(p)
}

/*SetNext is function */
func (d *Doctor) SetNext(next Department) {
	d.next = next
}