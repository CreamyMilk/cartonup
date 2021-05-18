package callback

import (
	"errors"
	"fmt"

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
	billSplit := c.BillRefNumber.Split("#")
	if len(billSplit) != 2 {
		return false
	}
	c.houseNo = billSplit[1]
	return c.BillRefNumber.HasPrefix("R#")
}

func (c *PCallback) isDeposit() bool {
	billSplit := c.BillRefNumber.Split("#")
	if len(billSplit) != 2 {
		return false
	}
	c.walletName = billSplit[1]
	return c.BillRefNumber.HasPrefix("DF#")
}
func (c *PCallback) classify() error {
	if c.isRentPayment() {
		//Get Tenants Details
		ten := tenant.GetTenantByHouseNo(c.houseNo)
		if ten == nil {
			return errors.New("The House must have been closed or is no longer in operation")
		}
		if ten.AmountDue != int64(c.TransAmount) {
			//Handle this as a deposit instead
			//Also rich fellas who decide to overpay we just deposit also
			return errors.New("So someone managed to send a malicous requst so ")
		}
		err := ten.ClearPayment()
		if err != nil {
			return err
		}
		//if we are here we need to notify freaking everyone so
		err = sms.SendSuccesfulPayment(c)
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
			return sms.SendNoWalletFound(c)
		}
		if err := wallet.Deposit(int64(c.TransAmount)); err != nil {
			return err
		}
		err = sms.SendWalletDepoistSuccess()
		if err != nil {
			fmt.Println(err)
		}
		//FCM messages are send by the wallet directly
	}
}
