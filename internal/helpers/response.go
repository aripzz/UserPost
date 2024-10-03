package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type StandardResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := StandardResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, status int, message string) error {
	response := StandardResponse{
		Status:  status,
		Message: message,
	}
	return c.Status(status).JSON(response)
}
