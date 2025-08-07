package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
	"go.uber.org/zap"
)

func FiberErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var appErr *errorhandler.AppError
		if errors.As(err, &appErr) {
			logger.Log.Error("Handled application error",
				zap.String("code", appErr.Code),
				zap.String("message", appErr.Message),
				zap.Int("status", appErr.StatusCode),
				zap.String("path", c.OriginalURL()),
			)
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
			logger.Log.Warn("Validation error", zap.Error(err))
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "VALIDATION_ERROR",
					"message": "Validation failed",
					"details": errors,
				},
			})
		}

		logger.Log.Error("Unhandled error",
			zap.Error(err),
			zap.String("path", c.OriginalURL()),
		)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    "UNKNOWN_ERROR",
				"message": err.Error(),
			},
		})
	}
}
