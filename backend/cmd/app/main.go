package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yusufbulac/byfood-case/backend/internal/repository"
	"github.com/yusufbulac/byfood-case/backend/internal/routes"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
	"github.com/yusufbulac/byfood-case/backend/pkg/config"
	"github.com/yusufbulac/byfood-case/backend/pkg/database"
	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
	"github.com/yusufbulac/byfood-case/backend/pkg/validator"
)

func main() {
	// Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	// Config, Validator, DB
	config.LoadConfig()
	db := database.ConnectMySQL(logger.Log)
	validator.InitValidator()

	// Fiber app
	app := fiber.New()

	// Book Module
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)

	// Routes
	v1 := app.Group("/api/v1")
	routes.RegisterBookRoutes(v1, bookService)

	// Root route
	app.Get("/", func(c fiber.Ctx) error {
		logger.Log.Info("Root endpoint hit")
		return c.SendString("Backend is running...")
	})

	app.Listen(":" + config.AppConfig.AppPort)
}
