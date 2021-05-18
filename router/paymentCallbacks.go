package router

import (
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
	return c.JSON(&fiber.Map{
		"status":     0,
		"chechoutID": req.TransID,
	})
}
