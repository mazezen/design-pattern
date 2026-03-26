package templatemethod

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	smsOtp := &Sms{}
	o := Otp{
		iOtp: smsOtp,
	}
	o.genAndSendOTP(4)

	fmt.Println()

	emailOtp := &Email{}
	o = Otp{
		iOtp: emailOtp,
	}
	o.genAndSendOTP(4)

}