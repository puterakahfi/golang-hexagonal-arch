package service

import (
	"golang-hexagonal-arch/module/book/dto"
	"golang-hexagonal-arch/module/book/entity"
)

func (s *bookService) RegisterNewBook(bookRequest dto.BookDto) (entity.Book, error) {
	newBook := dto.AssembleToEntity(bookRequest)
	books, err := s.repository.RegisterNewBook(newBook)
	return books, err
}
