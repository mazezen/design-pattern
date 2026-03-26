package factorymethod

type EmailSender struct{}

func (e *EmailSender) Send(msg string) string {
	return "Email: " + msg
}