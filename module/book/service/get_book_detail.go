package service

import "golang-hexagonal-arch/module/book/entity"

func (s *bookService) GetBookDetailId(id int) (entity.Book, error) {
	book, err := s.repository.GetBookDetailId(id)
	return book, err
}
