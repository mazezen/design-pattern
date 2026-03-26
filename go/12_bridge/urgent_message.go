package bridge

type UrgentMessage struct {
	sender Sender
}

func NewUrgentMessage(sender Sender) *UrgentMessage {
	return &UrgentMessage{sender: sender}
}

func (m *UrgentMessage) SendMessage(content string) string {
	return m.sender.Send("[Urgent] " + content)
}