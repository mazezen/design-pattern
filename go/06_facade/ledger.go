package facade

import "fmt"

// Ledger 外观模式组成部分 账簿
type Ledger struct {}

func (s *Ledger) makeEntry(accountName, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountName, txnType, amount)
    return
}