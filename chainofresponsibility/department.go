package chainofresponsibility

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}