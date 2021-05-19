package callback

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/CreamyMilk/cartonup/sms"
	"github.com/CreamyMilk/cartonup/tenant"
	"github.com/CreamyMilk/cartonup/wallet"
)

type PCallback struct {
	TransactionType   string `json:"TransactionType"`
	TransID           string `json:"TransID"`
	TransTime         string `json:"TransTime"`
	TransAmount       string `json:"TransAmount"`
	BusinessShortCode string `json:"BusinessShortCode"`
	BillRefNumber     string `json:"BillRefNumber"`
	InvoiceNumber     string `json:"InvoiceNumber"`
	OrgAccountBalance string `json:"OrgAccountBalance"`
	ThirdPartyTransID string `json:"ThirdPartyTransID"`
	MSISDN            string `json:"MSISDN"`
	FirstName         string `json:"FirstName"`
	MiddleName        string `json:"MiddleName"`
	LastName          string `json:"LastName"`
	houseNo           string
	walletName        string
}

func (c *PCallback) isRentPayment() bool {
	billSplit := strings.Split(c.BillRefNumber, "#")
	if len(billSplit) != 2 {
		return false
	}
	c.houseNo = billSplit[1]
	return strings.HasPrefix(c.BillRefNumber, "R#")
}

func (c *PCallback) isDeposit() bool {
	billSplit := strings.Split(c.BillRefNumber, "#")
	if len(billSplit) != 2 {
		return false
	}
	c.walletName = billSplit[1]
	return strings.HasPrefix(c.BillRefNumber, "DF#")
}

func (c *PCallback) Classify() error {
	if c.isRentPayment() {
		//Get Tenants Details
		ten := tenant.GetTenantByHouseNo(c.houseNo)
		if ten == nil {
			return errors.New("the House must have been closed or is no longer in operation")
		}
		transactionAmount, err := strconv.Atoi(c.TransAmount)
		if err != nil {
			fmt.Println(err)
		}
		if ten.AmountDue != (transactionAmount) {
			//Handle this as a deposit instead
			//Also rich fellas who decide to overpay we just deposit also
			return errors.New("so someone managed to send a malicous requst so ")
		}
		err = ten.ClearPayment()
		if err != nil {
			return err
		}
		//if we are here we need to notify freaking everyone so
		err = sms.SendSuccesfulPayment()
		if err != nil {
			//we can store some where in some form of retry que
			fmt.Println(err)
		}
		//We can notify the owners payment was done today
		return nil
	}

	if c.isDeposit() {
		userWallet := wallet.GetWalletByName(c.walletName)
		if userWallet == nil {
			return sms.SendNoWalletFound()
		}
		dAmount, err := strconv.Atoi(c.TransAmount)
		if err != nil {
			fmt.Println(err)
		}

		if userWallet.Deposit(int64(dAmount)) {
			return errors.New("error depositing funds")
		}
		err = sms.SendWalletDepositSuccess()
		if err != nil {
			fmt.Println(err)
		}
		//FCM messages are send by the wallet directly
	}
	return nil
}
