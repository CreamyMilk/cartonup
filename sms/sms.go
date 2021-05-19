package sms

import (
	"fmt"
)

func SendSuccesfulPayment() error {
	//SMS
	fmt.Print(`
Dear ,
Your transaction of Kshs.${TransAmount} has successfuly
been credited to ICRIB Account ${BillRefNumber} 
at ${TransTime}, Ref. Number ${TransID} Thank You.
`)
	return nil
}

func SendNoWalletFound() error {
	fmt.Print(`
So my customer hauna walllet kwa system 
`)
	return nil
}

func SendWalletDepositSuccess() error {
	fmt.Println(`You have deposiyed 394349 to johns account`)
	return nil
}
