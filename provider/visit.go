package provider

import (
	"github.com/gofiber/fiber/v2"
)

func init() {
	NewRegister.RegisterProvider(RegisterComponent{
		Title:       "Total visitor",
		Endpoint:    "/visitor",
		Method:      "GET",
		Description: "Mendapatkan token picsart",
		Params:      map[string]interface{}{},
		Type:        "",
		Body:        map[string]interface{}{},

		Code: func(c *fiber.Ctx) error {
			return c.Status(200).JSON(VS)
		},
	})
}
