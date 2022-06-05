package usecase

import (
	"golang-hexagonal-arch/module/book/domain"
	"golang-hexagonal-arch/module/book/domain/entity"
	"golang-hexagonal-arch/module/book/dto"
)

// Book service implementation(adapter)
type bookUseCase struct {
	repository domain.BookRepositoryInterface
}

func NewBookUseCase(r domain.BookRepositoryInterface) *bookUseCase {
	return &bookUseCase{r}
}

func (s *bookUseCase) GetRepository() domain.BookRepositoryInterface {
	return s.repository
}

func (s *bookUseCase) GetListBooks() ([]dto.BookResponseDto, error) {
	books, err := s.repository.GetListAllBooks()

	var booksResponse []dto.BookResponseDto

	for _, bookItem := range books {
		bookReponse := dto.AssembleEntityToResponse(&bookItem)
		booksResponse = append(booksResponse, bookReponse)

	}

	return booksResponse, err
}

func (s *bookUseCase) GetBookDetailId(id int) (entity.Book, error) {
	book, err := s.repository.GetBookDetailId(id)
	return book, err
}

func (s *bookUseCase) DeleteBook(id int) (entity.Book, error) {
	book, err := s.repository.DeleteBook(id)
	return book, err
}

func (s *bookUseCase) UpdateBook(id int) (entity.Book, error) {

	book, err := s.repository.UpdateBook(id)
	return book, err
}

func (s *bookUseCase) RegisterNewBook(bookRequest dto.BookRequestDto) (entity.Book, error) {
	newBook := dto.AssembleRequestToEntity(bookRequest)
	books, err := s.repository.RegisterNewBook(newBook)
	return books, err
}
