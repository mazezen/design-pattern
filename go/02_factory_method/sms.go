package factorymethod

type SmsSender struct{}

func (s *SmsSender) Send(msg string) string {
	return "SMS: " + msg
}