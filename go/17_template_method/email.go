package templatemethod

import "fmt"

type Email struct {
	Otp
}

func (e *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
    fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (e *Email)sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
    return nil
}