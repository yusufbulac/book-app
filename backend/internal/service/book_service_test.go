package service_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/mocks"
	"github.com/yusufbulac/byfood-case/backend/internal/model"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
)

func TestBookService_GetAll(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookSvc := service.NewBookService(mockRepo)

	mockBooks := []model.Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Year: 2020},
		{ID: 2, Title: "Book 2", Author: "Author 2", Year: 2021},
	}

	mockRepo.On("GetAll").Return(mockBooks, nil)

	books, err := bookSvc.GetAll()

	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, "Book 1", books[0].Title)
	assert.Equal(t, "Author 2", books[1].Author)
}

func TestBookService_GetByID_Success(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookSvc := service.NewBookService(mockRepo)

	book := &model.Book{ID: 1, Title: "Golang 101", Author: "John Doe", Year: 2023}
	mockRepo.On("GetByID", uint(1)).Return(book, nil)

	result, err := bookSvc.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Golang 101", result.Title)
}

func TestBookService_GetByID_NotFound(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookSvc := service.NewBookService(mockRepo)

	mockRepo.On("GetByID", uint(99)).Return(nil, errors.New("not found"))

	_, err := bookSvc.GetByID(99)

	assert.Error(t, err)
}

func TestBookService_Create(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookSvc := service.NewBookService(mockRepo)

	input := dto.CreateBookRequest{Title: "New Book", Author: "Yusuf", Year: 2023}

	mockRepo.On("Create", mock.MatchedBy(func(b *model.Book) bool {
		return b.Title == input.Title && b.Author == input.Author && b.Year == input.Year
	})).Return(nil)

	result, err := bookSvc.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, "New Book", result.Title)
}

func TestBookService_Update(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookSvc := service.NewBookService(mockRepo)

	id := uint(1)
	existing := &model.Book{ID: id, Title: "Old", Author: "Author", Year: 2022}
	updateReq := dto.UpdateBookRequest{
		Title:  strPtr("Updated"),
		Author: strPtr("Author"),
		Year:   intPtr(2023),
	}

	mockRepo.On("GetByID", id).Return(existing, nil)
	mockRepo.On("Update", mock.MatchedBy(func(b *model.Book) bool {
		return b.ID == id && b.Title == "Updated" && b.Year == 2023
	})).Return(nil)

	result, err := bookSvc.Update(id, updateReq)

	assert.NoError(t, err)
	assert.Equal(t, "Updated", result.Title)
	assert.Equal(t, 2023, result.Year)
}

func TestBookService_Delete(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	bookSvc := service.NewBookService(mockRepo)

	id := uint(1)

	mockRepo.On("Delete", id).Return(nil)

	err := bookSvc.Delete(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// --- Helpers ---

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
