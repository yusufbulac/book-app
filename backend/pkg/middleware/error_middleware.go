package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
)

func FiberErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var appErr *errorhandler.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.StatusCode).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := make([]fiber.Map, len(ve))
			for i, e := range ve {
				errors[i] = fiber.Map{
					"field": e.Field(),
					"error": "failed on " + e.Tag(),
				}
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "VALIDATION_ERROR",
					"message": "Validation failed",
					"details": errors,
				},
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    "UNKNOWN_ERROR",
				"message": err.Error(),
			},
		})
	}
}
