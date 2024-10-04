package middleware

import (
	"strings"

	"User-Post-Backend/infra/logger"
	"User-Post-Backend/internal/constant"
	"User-Post-Backend/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorHandlerMiddleware(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return HandleError(c, err)
	}
	return nil
}

func HandleError(c *fiber.Ctx, err error) error {
	if strings.Contains(err.Error(), constant.ErrForeignKey) {
		logger.Errorln(err)

		return c.Status(fiber.StatusBadRequest).JSON(&helpers.StandardResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Terjadi kesalahan data tidak sesuai (foreign key). Silakan periksa data yang Anda masukkan.",
		})
	}
	if strings.Contains(err.Error(), constant.ErrRecordNotFound) {
		logger.Errorln(err)
		return c.Status(fiber.StatusNotFound).JSON(&helpers.StandardResponse{
			Status:  fiber.StatusNotFound,
			Message: "Data tidak ditemukan. Silakan periksa ID yang Anda masukkan.",
		})
	}
	logger.Errorln(err)
	return c.Status(fiber.StatusInternalServerError).JSON(&helpers.StandardResponse{
		Status:  fiber.StatusInternalServerError,
		Message: "Terjadi kesalahan, silakan coba lagi.",
	})
}
