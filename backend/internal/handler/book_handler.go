package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/pkg/response"
	"strconv"

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

// GetAll godoc
// @Summary     Get all books
// @Description Returns a list of all books
// @Tags        books
// @Produce     json
// @Success     200 {array}  dto.BookResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /books [get]
func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	books, err := h.service.GetAll()
	if err != nil {
		return errorhandler.InternalError("Failed to fetch books")
	}
	return response.Success(c, books)
}

// GetByID godoc
// @Summary     Get book by ID
// @Description Returns a book by its ID
// @Tags        books
// @Produce     json
// @Param       id path int true "Book ID"
// @Success     200 {object} dto.BookResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     404 {object} response.ErrorResponse
// @Router      /books/{id} [get]
func (h *BookHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	book, err := h.service.GetByID(uint(id))
	if err != nil {
		return err
	}
	return response.Success(c, book)
}

// Create godoc
// @Summary     Create a new book
// @Description Creates a new book with the given payload
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       book body dto.CreateBookRequest true "Create Book"
// @Success     201 {object} dto.BookResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /books [post]
func (h *BookHandler) Create(c *fiber.Ctx) error {
	var input dto.CreateBookRequest
	if err := c.BodyParser(&input); err != nil {
		return errorhandler.BadRequest("INVALID_PAYLOAD", "Failed to parse request body")
	}
	if err := validator.ValidateStruct(input); err != nil {
		return err
	}
	book, err := h.service.Create(input)
	if err != nil {
		return err
	}
	return response.Created(c, book)
}

// Update godoc
// @Summary     Update a book
// @Description Updates an existing book by ID
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       id path int true "Book ID"
// @Param       book body dto.UpdateBookRequest true "Update Book"
// @Success     200 {object} dto.BookResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     404 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /books/{id} [put]
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	var input dto.UpdateBookRequest
	if err := c.BodyParser(&input); err != nil {
		return errorhandler.BadRequest("INVALID_PAYLOAD", "Failed to parse request body")
	}
	if err := validator.ValidateStruct(input); err != nil {
		return err
	}
	book, err := h.service.Update(uint(id), input)
	if err != nil {
		return err
	}
	return response.Success(c, book)
}

// Delete godoc
// @Summary     Delete a book
// @Description Deletes a book by its ID
// @Tags        books
// @Produce     json
// @Param       id path int true "Book ID"
// @Success     204 {string} string "No Content"
// @Failure     400 {object} response.ErrorResponse
// @Failure     404 {object} response.ErrorResponse
// @Router      /books/{id} [delete]
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorhandler.BadRequest("INVALID_ID", "ID must be a number")
	}
	if err := h.service.Delete(uint(id)); err != nil {
		return err
	}
	return response.NoContent(c)
}
