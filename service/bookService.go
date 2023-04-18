package service

import (
	"pustaka-api/entity"
	"pustaka-api/repository"
)

type BookService interface {
	GetBook() ([]entity.BookResponse, error)
	ShowBook(ID uint) (*entity.BookResponse, error)
	StoreBook(book entity.BookInput) (entity.Book, error)
	UpdateBook(ID uint, bookInput entity.BookInput) (entity.Book, error)
	DeleteBook(ID uint) error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{r}
}

func (s *bookService) GetBook() ([]entity.BookResponse, error) {
	book, err := s.bookRepository.FindAllBook()
	if err != nil {
		return nil, err
	}

	var returnBook []entity.BookResponse
	for _, val := range book {
		returnBook = append(returnBook, entity.BookResponse{
			Title: val.Title,
			Price: val.Price,
		})
	}

	return returnBook, err
}

func (s *bookService) ShowBook(ID uint) (*entity.BookResponse, error) {
	book, err := s.bookRepository.FindBookById(ID)
	if err != nil {
		return nil, err
	}

	bookResponse := entity.BookResponse{
		Title: book.Title,
		Price: book.Price,
	}

	return &bookResponse, err
}

func (s *bookService) StoreBook(book entity.BookInput) (entity.Book, error) {
	b, err := s.bookRepository.CreateBook(book)

	return b, err
}

func (s *bookService) UpdateBook(ID uint, book entity.BookInput) (entity.Book, error) {
	b, err := s.bookRepository.UpdateBook(ID, book)

	return b, err
}

func (s *bookService) DeleteBook(ID uint) error {
	return s.bookRepository.DeleteBook(ID)
}
