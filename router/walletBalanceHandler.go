package router

import (
	"github.com/CreamyMilk/cartonup/wallet"
	"github.com/gofiber/fiber/v2"
)

type balanceRequest struct {
	WalletName string `json:"walletname"`
}

func walletBalanceHandler(c *fiber.Ctx) error {
	req := new(balanceRequest)

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
	return c.JSON(&fiber.Map{
		"status":  0,
		"balance": w.GetBalance(),
	})
}
