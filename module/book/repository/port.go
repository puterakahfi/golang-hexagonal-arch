package repository

import (
	"golang-hexagonal-arch/module/book/entity"
)

type BookRepositoryInterface interface {
	GetListAllBooks() ([]entity.Book, error)
	RegisterNewBook(entity.Book) (entity.Book, error)
	GetBookDetailId(id int) (entity.Book, error)
	DeleteBook(id int) (entity.Book, error)
}
