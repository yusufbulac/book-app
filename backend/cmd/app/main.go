package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		logger.Log.Info("Root endpoint hit")
		return c.SendString("Backend is running...")
	})

	if err := app.Listen(":8080"); err != nil {
		logger.Log.Fatal("Server failed to start", zap.Error(err))
	}
}
