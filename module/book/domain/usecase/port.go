package usecase

import (
	"golang-hexagonal-arch/module/book/domain/entity"
	"golang-hexagonal-arch/module/book/dto"
)

// Book use cases  contract(port)
type BookUseCaseInterface interface {
	GetBookDetailId(id int) (entity.Book, error)
	RegisterNewBook(dto.BookRequestDto) (entity.Book, error)
	DeleteBook(id int) (entity.Book, error)
	UpdateBook(id int) (entity.Book, error)
	GetListBooks() ([]dto.BookResponseDto, error)
}
