package router

import (
	"github.com/CreamyMilk/cartonup/tenant"
	"github.com/CreamyMilk/cartonup/wallet"
	"github.com/gofiber/fiber/v2"
)

type WalletRequest struct {
	WalletName  string `json:"walletname"`
	HouseNumber string `json:"houseName"`
}

func walletPaymentHandler(c *fiber.Ctx) error {
	req := new(WalletRequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	w := wallet.GetWalletByName(req.WalletName)
	if w == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Seems The wallet provided is invalid",
		})
	}
	t := tenant.GetTenantByHouseNo(req.HouseNumber)
	if t == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Could not retrive your payment details",
		})
	}
	if w.GetBalance() < t.AmountDue {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Seems you need to deposit more funds to complete payment",
		})
	}
	superWallet := wallet.GetRentReceiverWallet()
	if superWallet == nil {
		return c.JSON(&fiber.Map{
			"status":  -4,
			"message": "Currently under maintaince try again later",
		})
	}
	if t.RentStatus != 0 {
		t.ClearPayment()
		return c.JSON(&fiber.Map{
			"status":  0,
			"message": "Seem you have no due payments",
		})
	}
	why, success := w.SendMoney(int64(t.AmountDue), *superWallet)
	if !success {
		return c.JSON(&fiber.Map{
			"status":  -5,
			"message": why,
		})
	}
	t.ClearPayment()
	return c.JSON(&fiber.Map{
		"status":  0,
		"message": "Rent Payment Was successful",
	})

}
