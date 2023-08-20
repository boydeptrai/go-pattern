package chainofresponsibility

import (
	"fmt"
)
/*Cashier is struct*/
type Cashier struct {
	next Department
}
/*Execute is fuction*/
func (c *Cashier) Execute(p *Patient) {
	if p.isPaid {
		fmt.Println("Patient already paid their bill")
		return
	}
	fmt.Println("Patient is paying the bill")
	p.isPaid = true
}

func (c *Cashier) SetNext (next Department) {
	c.next = next
}