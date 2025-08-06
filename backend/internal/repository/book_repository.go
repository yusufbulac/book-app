package repository

import (
	"github.com/yusufbulac/byfood-case/backend/internal/model"
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
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}
