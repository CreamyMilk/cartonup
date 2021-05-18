package wallet

type wallet struct {
	name    string
	balance int64
}

func GetWalletByName(walletName string) *Wallet {
	return nil
}
func (w *Wallet) Deposit(amount int64) error {
	//Send FCM notifcation to that effect
	return nil
}
