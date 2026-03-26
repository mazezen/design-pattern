package mediator

import "testing"

func TestChatRoom(t *testing.T) {
	room := NewChatRoom()

	u1 := NewUser("A", room)
	u2 := NewUser("B", room)
	u3 := NewUser("C", room)

	room.Register(u1)
	room.Register(u2)
	room.Register(u3)

	u1.Send("hello")

	// u2 和 u3 应该收到
	if len(u2.Messages) != 1 || u2.Messages[0] != "A: hello" {
		t.Errorf("u2 did not receive message")
	}

	if len(u3.Messages) != 1 || u3.Messages[0] != "A: hello" {
		t.Errorf("u3 did not receive message")
	}

	// u1 不应该收到自己的消息
	if len(u1.Messages) != 0 {
		t.Errorf("u1 should not receive its own message")
	}
}