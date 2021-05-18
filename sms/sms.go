package sms

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/callback"
)

func SendSuccesfulPayment(p *callback.PCallback) error {
	//SMS
	fmt.Println(`
Dear %s %s,
Your transaction of Kshs.${TransAmount} has successfuly
been credited to ICRIB Account ${BillRefNumber} 
at ${TransTime}, Ref. Number ${TransID} Thank You.
`)
	return nil
}

func SendNoWalletFound(p *callback.PCallback) error {
	fmt.Println(`
So my customer hauna walllet kwa system 
`)
}

func SendWalletDepositSuccess() error {
	fmt.Println(`You have deposiyed 394349 to johns account`)

}
