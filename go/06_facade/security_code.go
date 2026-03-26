package facade

import "fmt"

// SecurityCode 外观模式组成部分 安全码
type SecurityCode struct {
	code string
}

func newSecurityCode(code string) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(incomingCode string) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security code is incorrect")
	}
	fmt.Println("SecurityCode verified")

	return nil
}