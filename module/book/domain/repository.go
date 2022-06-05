package domain

import (
	"golang-hexagonal-arch/module/book/domain/entity"
)

// Book Repository Contract (Port)
type BookRepositoryInterface interface {
	GetListAllBooks() ([]entity.Book, error)
	RegisterNewBook(entity.Book) (entity.Book, error)
	GetBookDetailId(id int) (entity.Book, error)
	DeleteBook(id int) (entity.Book, error)
	UpdateBook(id int) (entity.Book, error)
}
