package router

import (
	"fmt"
	"strconv"

	"github.com/CreamyMilk/cartonup/mpeza"
	"github.com/CreamyMilk/cartonup/wallet"
	"github.com/gofiber/fiber/v2"
)

type depositRequest struct {
	Wallet_name string `json:"walletname"`
	Phone       string `json:"phone"`
	Amount      int    `json:"amount"`
}

func depositRequestHandler(c *fiber.Ctx) error {
	req := new(depositRequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	w := wallet.GetWalletByName(req.Wallet_name)
	if w == nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": "Seems you lack an ICRIB wallet",
		})
	}
	checkoutID, err := mpeza.SendSTK(req.Phone, strconv.Itoa(req.Amount), "DF#"+req.Wallet_name, req.Wallet_name)
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": err.Error(),
		})
	}
	//Store this somewhere
	fmt.Print(checkoutID)
	return c.JSON(&fiber.Map{
		"status":     0,
		"chechoutID": checkoutID,
	})
}
