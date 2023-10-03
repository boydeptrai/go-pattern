package mediator

import "fmt"

type GoodTrain struct {
	Mediator Mediator
}

func (g *GoodTrain) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Println("Good Train: Landing")
	} else {
		fmt.Println("Good Train: Waiting")
	}
}

func (g *GoodTrain) Departure() {
	g.Mediator.notifyFree()
	fmt.Println("GoodTrain: Leaving")
}

func (g *GoodTrain) PermitArrival() {
	fmt.Println("Goodtrain: Arrival Permitted. Landing")
}