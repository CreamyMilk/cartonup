package wallet

import "github.com/CreamyMilk/cartonup/database"

type Wallet struct {
	name    string
	balance int64
}

func GetWalletByName(name string) *Wallet {
	tempWall := new(Wallet)
	err := database.DB.QueryRow("SELECT wallet_name,balance FROM wallets_store WHERE wallet_name=? ", name).Scan(&tempWall.name, &tempWall.balance)
	if err != nil {
		return nil
	}
	return tempWall
}

//Deposit is a transactional representaition of old deposit
func (w *Wallet) Deposit(amount int64) bool {
	tempBalance := 0
	tx, err := database.DB.Begin()
	if err != nil {
		tx.Rollback()
		return false
	}

	getBalStm, err := tx.Prepare("SELECT balance FROM wallets_store WHERE wallet_name = ?")
	getBalStm.QueryRow(w.name).Scan(&tempBalance)
	defer getBalStm.Close()
	if err != nil {
		return false
	}
	newBalance := int64(tempBalance) + amount
	if amount < 1 {
		tx.Rollback()
		return false
	}
	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newBalance, w.name)
	if err != nil {
		tx.Rollback()
		//fmt.Printf("-------------->%v", err)
	}
	w.balance = int64(tempBalance)
	tx.Commit()
	return true
}
