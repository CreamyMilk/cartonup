package router

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/callback"
	"github.com/gofiber/fiber/v2"
)

func successfulPaymentHandler(c *fiber.Ctx) error {
	req := new(callback.PCallback)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}
	err := req.Classify()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(&fiber.Map{
		"status":     0,
		"chechoutID": req.TransID,
	})
}
