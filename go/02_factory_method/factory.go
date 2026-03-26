package factorymethod

type Factory interface {
	Create() Sender
}