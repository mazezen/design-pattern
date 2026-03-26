package mediator

type Mediator interface {
	Send(msg string, sender Colleague)
}

