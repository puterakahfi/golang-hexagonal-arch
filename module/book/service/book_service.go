package service

import (
	"golang-hexagonal-arch/module/book/dto"
	"golang-hexagonal-arch/module/book/entity"
	"golang-hexagonal-arch/module/book/repository"
)

type bookService struct {
	repository repository.BookRepositoryInterface
}

func NewService(r repository.BookRepositoryInterface) *bookService {
	return &bookService{r}
}

func (s *bookService) GetRepository() repository.BookRepositoryInterface {
	return s.repository
}

func (s *bookService) GetListBooks() ([]entity.Book, error) {
	books, err := s.repository.GetListAllBooks()
	return books, err
}

func (s *bookService) GetBookDetailId(id int) (entity.Book, error) {
	book, err := s.repository.GetBookDetailId(id)
	return book, err
}

func (s *bookService) DeleteBook(id int) (entity.Book, error) {
	book, err := s.repository.DeleteBook(id)
	return book, err
}

func (s *bookService) RegisterNewBook(bookRequest dto.BookDto) (entity.Book, error) {
	newBook := dto.AssembleToEntity(bookRequest)
	books, err := s.repository.RegisterNewBook(newBook)
	return books, err
}
