package service

import "pustaka-api/repository"

type Service struct {
	Book BookService 
}

func Init(repository *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repository.Book),
	}
}