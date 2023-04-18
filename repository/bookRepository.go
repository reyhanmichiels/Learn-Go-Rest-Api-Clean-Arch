package repository

import (
	"pustaka-api/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAllBook() ([]entity.Book, error)
	FindBookById(ID uint) (entity.Book, error)
	CreateBook(book entity.BookInput) (entity.Book, error)
	UpdateBook(ID uint, book entity.BookInput) (entity.Book, error)
	DeleteBook(ID uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) FindAllBook() ([]entity.Book, error) {
	var book []entity.Book

	err := r.db.Find(&book).Error
	if err != nil {
		return nil, err
	}

	return book, err
}

func (r *bookRepository) FindBookById(ID uint) (entity.Book, error) {
	var book entity.Book

	err := r.db.First(&book, ID).Error
	// err := r.db.Where("id = ?", ID).First(&book).Error
	return book, err
}

func (r *bookRepository) CreateBook(book entity.BookInput) (entity.Book, error) {
	b := entity.Book{
		Title: book.Title,
		Price: book.Price,
	}

	err := r.db.Create(&b).Error

	return b, err
}

func (r *bookRepository) UpdateBook(ID uint, book entity.BookInput) (entity.Book, error) {
	b, err := r.FindBookById(ID)
	if err != nil {
		return b, err
	}

	b.Title = book.Title
	b.Price = book.Price

	err = r.db.Save(&b).Error
	return b, err
}

func (r *bookRepository) DeleteBook(ID uint) error {
	return r.db.Delete(&entity.Book{}, ID).Error
}
