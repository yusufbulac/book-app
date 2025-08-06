package repository

import (
	"errors"

	"github.com/yusufbulac/byfood-case/backend/internal/model"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetByID(id uint) (*model.Book, error)
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, errorhandler.InternalError("Failed to get books")
	}
	return books, nil
}

func (r *bookRepository) GetByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorhandler.NotFound("BOOK_NOT_FOUND", "Book not found")
	}
	if err != nil {
		return nil, errorhandler.InternalError("Failed to get book by ID")
	}
	return &book, nil
}

func (r *bookRepository) Create(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return errorhandler.InternalError("Failed to create book")
	}
	return nil
}

func (r *bookRepository) Update(book *model.Book) error {
	if err := r.db.Save(book).Error; err != nil {
		return errorhandler.InternalError("Failed to update book")
	}
	return nil
}

func (r *bookRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Book{}, id)
	if result.Error != nil {
		return errorhandler.InternalError("Failed to delete book")
	}
	if result.RowsAffected == 0 {
		return errorhandler.NotFound("BOOK_NOT_FOUND", "Book not found")
	}
	return nil
}
