package bridge

type EmailSender struct{}

func (e *EmailSender) Send(content string) string {
	return "Email: " + content
}