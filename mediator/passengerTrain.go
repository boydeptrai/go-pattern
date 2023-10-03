package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (g *PassengerTrain) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *PassengerTrain) Departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyFree()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}