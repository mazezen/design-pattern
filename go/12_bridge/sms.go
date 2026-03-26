package bridge

type SmsSender struct{}

func (s *SmsSender) Send(content string) string {
	return "SMS: " + content
}