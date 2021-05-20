package router

import (
	"fmt"
	"strconv"

	"github.com/CreamyMilk/cartonup/mpeza"
	"github.com/CreamyMilk/cartonup/tenant"
	"github.com/gofiber/fiber/v2"
)

type paymentRequest struct {
	TenantID int    `json:"tenantID"`
	Phone    string `json:"phone"`
}

func paymentRequestHandler(c *fiber.Ctx) error {
	req := new(paymentRequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	tenant := tenant.GetTenantByID(req.TenantID)
	if tenant == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Error retriving your account details.",
		})
	}

	checkoutID, err := mpeza.SendSTK(req.Phone, strconv.Itoa(tenant.AmountDue), "R#"+tenant.HouseName, strconv.Itoa(req.TenantID))
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
