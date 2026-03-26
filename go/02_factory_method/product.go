package factorymethod

type Sender interface {
	Send(msg string) string
}