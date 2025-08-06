package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
	"github.com/yusufbulac/byfood-case/backend/pkg/validator"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetAll(c fiber.Ctx) error {
	books, err := h.service.GetAll()
	if err != nil {
		return errorhandler.InternalError("Failed to fetch books")
	}
	return c.JSON(books)
}

func (h *BookHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	book, err := h.service.GetByID(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(book)
}

func (h *BookHandler) Create(c fiber.Ctx) error {
	var input dto.CreateBookRequest
	if err := c.Bind().Body(&input); err != nil {
		return errorhandler.BadRequest("INVALID_PAYLOAD", "Failed to parse request body")
	}
	if err := validator.ValidateStruct(input); err != nil {
		return err
	}
	book, err := h.service.Create(input)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *BookHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	var input dto.UpdateBookRequest
	if err := c.Bind().Body(&input); err != nil {
		return errorhandler.BadRequest("INVALID_PAYLOAD", "Failed to parse request body")
	}
	if err := validator.ValidateStruct(input); err != nil {
		return err
	}
	book, err := h.service.Update(uint(id), input)
	if err != nil {
		return err
	}
	return c.JSON(book)
}

func (h *BookHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	if err := h.service.Delete(uint(id)); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
