package dto

import "golang-hexagonal-arch/module/book/domain/entity"

func AssembleRequestToEntity(dto BookRequestDto) entity.Book {

	price, _ := dto.Price.Int64()
	newBook := entity.Book{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       int(price),
	}
	return newBook
}

func AssembleEntityToResponse(book *entity.Book) BookResponseDto {

	bookResponse := BookResponseDto{
		Title:       book.Title,
		SubTitle:    book.SubTitle,
		Description: book.Description,
		Discount:    book.Discount,
	}
	return bookResponse
}
