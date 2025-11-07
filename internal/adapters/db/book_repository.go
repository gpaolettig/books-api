package db

import (
	"books-api/internal/core/book"
	"books-api/internal/core/shared"
	"errors"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) book.BookRepository {
	return &BookRepository{db: db}
}
func (r *BookRepository) FindById(id int) (book.Book, error) { //should i return a pointer?
	var bk book.Book
	if err := r.db.First(&bk, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return book.Book{}, shared.ErrBookNotFound
		}
		return book.Book{}, err
	}
	return bk, nil
}
func (r *BookRepository) FindAll() ([]book.Book, error) {
	var books []book.Book
	if err := r.db.Find(&books).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrNoBooksFound
		}
		return nil, err
	}
	return books, nil
}
