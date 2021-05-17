package router

import (
	"github.com/CreamyMilk/cartonup/mpeza"
	"github.com/CreamyMilk/cartonup/tenant"
	"github.com/gofiber/fiber/v2"
)

type paymentRequest struct {
	tenantID int    `json:"tenantID"`
	phone    string `json:"phone"`
}

func paymentRequestHandler(c *fiber.Ctx) error {
	req := new(paymentRequest)

	if err := c.BodyParser(req); err != nil {
		return c.JSON(&fiber.Map{
			"status":  -1,
			"message": "request is malformed",
		})
	}

	tenant := tenant.GetTenantByID(req.tenantID)
	if tenant == nil {
		return c.JSON(&fiber.Map{
			"status":  -2,
			"message": "Error retriving your accout details.",
		})
	}

	//mpesa.SendSTK(req.Mobile
	checkoutID, err := mpeza.SendSTK(req.phone, string(tenant.AmountDue), tenant.HouseName, string(req.tenantID))
	if err != nil {
		return c.JSON(&fiber.Map{
			"status":  -3,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"status":     0,
		"chechoutID": checkoutID,
	})
}
