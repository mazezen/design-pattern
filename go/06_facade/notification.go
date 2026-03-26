package facade

import "fmt"

// Nitification 外观模式组成部分 通知
type Nitification struct {}

func (n *Nitification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Nitification) sendWalletDebitNotification() {
    fmt.Println("Sending wallet debit notification")
}