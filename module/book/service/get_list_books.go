package service

import "golang-hexagonal-arch/module/book/entity"

func (s *bookService) GetListBooks() ([]entity.Book, error) {
	books, err := s.repository.GetListAllBooks()
	return books, err
}
