package service

import (
	"golang-hexagonal-arch/module/book/dto"
	"golang-hexagonal-arch/module/book/entity"
	"golang-hexagonal-arch/module/book/repository"
)

type BookServiceInterface interface {
	GetListBooks() ([]entity.Book, error)
	GetBookDetailId(id int) (entity.Book, error)
	RegisterNewBook(dto.BookDto) (entity.Book, error)
}

type bookService struct {
	repository repository.BookRepositoryInterface
}

func NewService(r repository.BookRepositoryInterface) *bookService {
	return &bookService{r}
}

func (s *bookService) GetRepository() repository.BookRepositoryInterface {
	return s.repository
}
