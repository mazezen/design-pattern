// bridge/bridge_test.go
package bridge

import "testing"

func TestCommonMessage(t *testing.T) {
	email := &EmailSender{}
	msg := NewCommonMessage(email)

	res := msg.SendMessage("test")

	if res != "Email: [Common] test" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestUrgentMessage(t *testing.T) {
	sms := &SmsSender{}
	msg := NewUrgentMessage(sms)

	res := msg.SendMessage("alert")

	if res != "SMS: [Urgent] alert" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestSwitchSender(t *testing.T) {
	email := &EmailSender{}
	sms := &SmsSender{}

	msg := NewUrgentMessage(email)
	res1 := msg.SendMessage("x")

	msg = NewUrgentMessage(sms)
	res2 := msg.SendMessage("x")

	if res1 != "Email: [Urgent] x" || res2 != "SMS: [Urgent] x" {
		t.Errorf("bridge failed")
	}
}