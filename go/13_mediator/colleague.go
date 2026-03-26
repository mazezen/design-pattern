package mediator

type Colleague interface {
	Receive(msg string)
}