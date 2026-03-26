package factorymethod

import "testing"

func TestSmsFactory(t *testing.T) {
	f := &SmsFactory{}
	s := f.Create()

	res := s.Send("test")

	if res != "SMS: test" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestEmailFactory(t *testing.T) {
	f := &EmailFactory{}
	s := f.Create()

	res := s.Send("test")

	if res != "Email: test" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestPolymorphism(t *testing.T) {
	var f Factory

	f = &SmsFactory{}
	if f.Create().Send("x") != "SMS: x" {
		t.Error("sms failed")
	}

	f = &EmailFactory{}
	if f.Create().Send("x") != "Email: x" {
		t.Error("email failed")
	}
}