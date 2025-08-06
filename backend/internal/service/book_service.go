package service

import (
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/model"
	"github.com/yusufbulac/byfood-case/backend/internal/repository"
)

type BookService interface {
	GetAll() ([]dto.BookResponse, error)
	GetByID(id uint) (*dto.BookResponse, error)
	Create(input dto.CreateBookRequest) (*dto.BookResponse, error)
	Update(id uint, input dto.UpdateBookRequest) (*dto.BookResponse, error)
	Delete(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAll() ([]dto.BookResponse, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var response []dto.BookResponse
	for _, book := range books {
		response = append(response, toBookResponse(book))
	}
	return response, nil
}

func (s *bookService) GetByID(id uint) (*dto.BookResponse, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	resp := toBookResponse(*book)
	return &resp, nil
}

func (s *bookService) Create(input dto.CreateBookRequest) (*dto.BookResponse, error) {
	book := model.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
	}
	if err := s.repo.Create(&book); err != nil {
		return nil, err
	}
	resp := toBookResponse(book)
	return &resp, nil
}

func (s *bookService) Update(id uint, input dto.UpdateBookRequest) (*dto.BookResponse, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Title != "" {
		book.Title = input.Title
	}
	if input.Author != "" {
		book.Author = input.Author
	}
	if input.Year != 0 {
		book.Year = input.Year
	}
	if err := s.repo.Update(book); err != nil {
		return nil, err
	}
	resp := toBookResponse(*book)
	return &resp, nil
}

func (s *bookService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func toBookResponse(book model.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Year:      book.Year,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}
