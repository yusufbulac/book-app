package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/mocks"
	"github.com/yusufbulac/byfood-case/backend/internal/model"
	"github.com/yusufbulac/byfood-case/backend/internal/routes"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
	"github.com/yusufbulac/byfood-case/backend/pkg/middleware"
	myvalidator "github.com/yusufbulac/byfood-case/backend/pkg/validator"
)

type successResp struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type errorResp struct {
	Success bool `json:"success"`
	Error   struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

func setupBooksApp(svc service.BookService) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.FiberErrorHandler(),
	})
	v1 := app.Group("/api/v1")
	routes.RegisterBookRoutes(v1, svc)
	return app
}

func TestBooks_GetAll_Success(t *testing.T) {
	myvalidator.InitValidator()

	mockRepo := mocks.NewBookRepository(t)
	svc := service.NewBookService(mockRepo)
	app := setupBooksApp(svc)

	mockBooks := []model.Book{
		{ID: 1, Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
		{ID: 2, Title: "Domain-Driven Design", Author: "Eric Evans", Year: 2003},
	}
	mockRepo.On("GetAll").Return(mockBooks, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var ok successResp
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&ok))

	var list []dto.BookResponse
	assert.NoError(t, json.Unmarshal(ok.Data, &list))
	assert.Len(t, list, 2)
	assert.Equal(t, "Clean Code", list[0].Title)
}

func TestBooks_GetByID_NotFound(t *testing.T) {
	myvalidator.InitValidator()

	mockRepo := mocks.NewBookRepository(t)
	svc := service.NewBookService(mockRepo)
	app := setupBooksApp(svc)

	mockRepo.On("GetByID", uint(999)).
		Return(nil, errorhandler.NotFound("BOOK_NOT_FOUND", "Book not found"))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/books/999", nil)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	var er errorResp
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&er))
	assert.Equal(t, "BOOK_NOT_FOUND", er.Error.Code)
}

func TestBooks_Create_ValidationError(t *testing.T) {
	myvalidator.InitValidator()

	mockRepo := mocks.NewBookRepository(t)
	svc := service.NewBookService(mockRepo)
	app := setupBooksApp(svc)

	// no title -> validation error
	payload := map[string]interface{}{
		"author": "Yusuf",
		"year":   2023,
	}
	b, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/books/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	var er errorResp
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&er))
	assert.Equal(t, "VALIDATION_ERROR", er.Error.Code)
}

func TestBooks_Create_Success(t *testing.T) {
	myvalidator.InitValidator()

	mockRepo := mocks.NewBookRepository(t)
	svc := service.NewBookService(mockRepo)
	app := setupBooksApp(svc)

	input := dto.CreateBookRequest{
		Title:  "New Book",
		Author: "Yusuf",
		Year:   2024,
	}

	mockRepo.On("Create", mock.MatchedBy(func(b *model.Book) bool {
		return b.Title == input.Title && b.Author == input.Author && b.Year == input.Year
	})).Return(nil)

	b, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var ok successResp
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&ok))

	var created dto.BookResponse
	assert.NoError(t, json.Unmarshal(ok.Data, &created))
	assert.Equal(t, "New Book", created.Title)
}

func TestBooks_Delete_Success(t *testing.T) {
	myvalidator.InitValidator()

	mockRepo := mocks.NewBookRepository(t)
	svc := service.NewBookService(mockRepo)
	app := setupBooksApp(svc)

	mockRepo.On("Delete", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/books/1", nil)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}
