package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yusufbulac/byfood-case/backend/internal/handler"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
)

func RegisterBookRoutes(router fiber.Router, service service.BookService) {
	h := handler.NewBookHandler(service)

	books := router.Group("/books")
	books.Get("/", h.GetAll)
	books.Get("/:id", h.GetByID)
	books.Post("/", h.Create)
	books.Put("/:id", h.Update)
	books.Delete("/:id", h.Delete)
}
