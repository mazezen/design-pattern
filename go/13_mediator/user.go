package mediator

type User struct {
	Name     string
	room     Mediator
	Messages []string // 用于测试记录
}

func NewUser(name string, room Mediator) *User {
	return &User{
		Name: name,
		room: room,
	}
}

func (u *User) Send(msg string) {
	u.room.Send(u.Name+": "+msg, u)
}

func (u *User) Receive(msg string) {
	u.Messages = append(u.Messages, msg)
}