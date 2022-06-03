package dto

import "golang-hexagonal-arch/module/book/entity"

func AssembleToEntity(dto BookDto) entity.Book {

	price, _ := dto.Price.Int64()
	newBook := entity.Book{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       int(price),
	}
	return newBook
}
