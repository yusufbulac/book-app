package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Proceed to next handler
		err := c.Next()

		// Log request after response is written
		logger.Log.Info("Incoming request",
			zap.String("method", c.Method()),
			zap.String("path", c.OriginalURL()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
			zap.String("ip", c.IP()),
			zap.String("user-agent", c.Get("User-Agent")),
		)

		return err
	}
}
