package facade

import "fmt"

// Account 外观模式组成部分 账户
type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{name: name}
}

func (a *Account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("Account name is incorrect")
	}
	fmt.Println("Account Verified")

	return nil
}