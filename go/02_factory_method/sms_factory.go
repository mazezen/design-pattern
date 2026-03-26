package factorymethod

type SmsFactory struct{}

func (f *SmsFactory) Create() Sender {
	return &SmsSender{}
}