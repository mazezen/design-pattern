package bridge

type CommonMessage struct {
	sender Sender
}

func NewCommonMessage(sender Sender) *CommonMessage {
	return &CommonMessage{sender: sender}
}

func (m *CommonMessage) SendMessage(content string) string {
	return m.sender.Send("[Common] " + content)
}