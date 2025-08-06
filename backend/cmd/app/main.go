package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yusufbulac/byfood-case/backend/pkg/config"
	"github.com/yusufbulac/byfood-case/backend/pkg/database"
	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	config.LoadConfig()
	database.ConnectMySQL(logger.Log)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		ctx := c.(fiber.Ctx)
		logger.Log.Info("Root endpoint hit")
		return ctx.SendString("Backend is running...")
	})

	app.Listen(":" + config.AppConfig.AppPort)
}
