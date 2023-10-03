package mediator

type Mediator interface {
	canLand(Train) bool
	notifyFree()
}