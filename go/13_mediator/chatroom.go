package mediator

type ChatRoom struct {
	users []Colleague
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) Register(user Colleague) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) Send(msg string, sender Colleague) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(msg)
		}
	}
}