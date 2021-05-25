package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/", homeHandler)

	v2 := app.Group("/v2")

	v2.Post("/paystk", paymentRequestHandler)
	v2.Post("/stkpushcall", stkcallbacksHandler)
	v2.Post("/successcall", successfulPaymentHandler)

	walletsGroup := v2.Group("/wallet")

	walletsGroup.Get("/balance", walletBalanceHandler)
	walletsGroup.Post("/pay", walletPaymentHandler)
	walletsGroup.Post("/deposit", depositRequestHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(418).JSON(&fiber.Map{
			"Message": "🍏 Route not found",
		}) // => 418 "I am a tepot"
	})
}

func homeHandler(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"Message": "Hello Handler",
	})
}
