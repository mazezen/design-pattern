package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	fmt.Println()
	walletFacade := newWalletFacade("tom", "124568");
	fmt.Println()

	err := walletFacade.addMoneyToWallet("tom", "124568", 10)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("tom", "124568", 5)
	if err != nil {
		t.Fatal(err)
	}
}