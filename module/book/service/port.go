package service

import (
	"golang-hexagonal-arch/module/book/dto"
	"golang-hexagonal-arch/module/book/entity"
)

type BookServiceInterface interface {
	GetListBooks() ([]entity.Book, error)
	GetBookDetailId(id int) (entity.Book, error)
	RegisterNewBook(dto.BookDto) (entity.Book, error)
	DeleteBook(id int) (entity.Book, error)
}
