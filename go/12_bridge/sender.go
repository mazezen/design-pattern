package bridge

type Sender interface {
	Send(content string) string
}