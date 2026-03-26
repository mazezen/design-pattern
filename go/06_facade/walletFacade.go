package facade

import "fmt"

// WalletFacade 外观模式
type WalletFacade struct {
	account *Account
	wallet *Wallet
	securityCode *SecurityCode
	nitification *Nitification
	ledger *Ledger
}

func newWalletFacade(accountName, code string) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		account: newAccount(accountName),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		nitification: &Nitification{},
		ledger: &Ledger{},
	}

	fmt.Println("Account created")
    return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(name, code string, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(name)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.nitification.sendWalletCreditNotification()
	w.ledger.makeEntry(name, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(name, code string, amount int) error {
	fmt.Println("Starting debit money from wallet")
	if err := w.account.checkAccount(name); err != nil {
		return err
	}
	err := w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.nitification.sendWalletDebitNotification()
	w.ledger.makeEntry(name, "debit", amount)

	return nil
}