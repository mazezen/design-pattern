package bridge

type Message interface {
	SendMessage(content string) string
}