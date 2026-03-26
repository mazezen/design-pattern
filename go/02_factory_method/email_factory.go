package factorymethod

type EmailFactory struct{}

func (f *EmailFactory) Create() Sender {
	return &EmailSender{}
}