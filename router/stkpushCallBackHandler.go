package router

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/notification"
	"github.com/gofiber/fiber/v2"
)

func stkcallbacksHandler(c *fiber.Ctx) error {
	r := new(notification.StkPushCallback)

	if err := c.BodyParser(r); err != nil {
		fmt.Println("Hello Enginnering this request structure has suddenly changed")
		return c.JSON(&fiber.Map{
			"ResponseCode": "00000000",
			"ResponseDesc": "success",
		})
	}
	err := r.Classify()
	if err != nil {
		fmt.Print(err)
	}
	return c.JSON(&fiber.Map{
		"ResponseCode": "00000000",
		"ResponseDesc": "success",
	})
}
