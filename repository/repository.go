package repository

import "gorm.io/gorm"

type Repository struct {
	Book BookRepository
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
	}
}
