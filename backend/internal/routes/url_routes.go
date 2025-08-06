package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/internal/handler"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
)

func RegisterUrlRoutes(router fiber.Router, urlService service.UrlService) {
	h := handler.NewUrlHandler(urlService)
	router.Post("/url/transform", h.Transform)
}
