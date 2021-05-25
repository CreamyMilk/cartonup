package wallet

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/database"
	"github.com/CreamyMilk/cartonup/notification"
)

func GetRentReceiverWallet() *Wallet {
	tempWall := new(Wallet)
	err := database.DB.QueryRow("SELECT wallet_name,balance FROM wallets_store WHERE wallet_name=? ", RENTRECEIVERNAME).Scan(&tempWall.name, &tempWall.balance)
	if err != nil {
		return nil
	}
	return tempWall
}

//SendMoney is used to move money from account a to account b
func (w *Wallet) SendMoney(amountToSend int64, recipientW Wallet) (string, bool) {
	transactionCost := 0
	tx, err := database.DB.Begin()
	errorMessage := ""
	if err != nil {
		tx.Rollback()
		return errorMessage, false
	}
	if amountToSend <= 0 {
		errorMessage = "You cannot send negative values."
		tx.Rollback()
		return errorMessage, false
	}
	//You cannot send to self
	if w.name == recipientW.name {
		errorMessage = "Cannot send funds to self"
		tx.Rollback()
		return errorMessage, false
	}
	//get current balance
	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", w.name).Scan(&w.balance)
	if err != nil {
		errorMessage = "Could not retrieve your balance."
		tx.Rollback()
		return errorMessage, false
	}

	if amountToSend >= w.balance {
		errorMessage = "Does not have transaction cost."
		tx.Rollback()
		return errorMessage, false
	}

	if recipientW.name == "" {
		errorMessage = "The Receipient seems to be invalid."
		tx.Rollback()
		return errorMessage, false
	}

	newSenderBalance := w.balance - (amountToSend)
	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newSenderBalance, w.name)
	if err != nil {
		errorMessage = "No update to sender balance"
		tx.Rollback()
		return errorMessage, false
	}

	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", recipientW.name).Scan(&recipientW.balance)
	if err != nil {
		errorMessage = "Could not get current balance"
		tx.Rollback()
		return errorMessage, false
	}
	newRecipientBalance := recipientW.balance + amountToSend

	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newRecipientBalance, recipientW.name)
	if err != nil {
		errorMessage = "No update to receiver balance"
		tx.Rollback()
		return errorMessage, false
	}

	//generate transaction of what had occured
	genereatedId := uuidgen()
	_, err = tx.Exec("INSERT INTO transactions_list (transuuid,sender_name,receiver_name,amount,charge,ttype) VALUES (?,?,?,?,?,?)", genereatedId, w.name, recipientW.name, amountToSend, transactionCost, SENDMONEY_TYPE)
	if err != nil {
		errorMessage = ""
		tx.Rollback()
		return errorMessage, false
	}
	//Applys Changes to the database
	tx.Commit()

	_, err = notification.SendNotification(w.name, notification.SENDING_MONEY, amountToSend)
	if err != nil {
		fmt.Printf("Failed to send notifcation because %v", err)
	}
	_, err = notification.SendNotification(recipientW.name, notification.RECEVIEING_MONEY, amountToSend)
	if err != nil {
		fmt.Printf("Failed to send notifcation because %v", err)
	}
	return "", true
}
