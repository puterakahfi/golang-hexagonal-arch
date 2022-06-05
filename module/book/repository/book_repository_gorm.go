package repository

// Book Repository GORM Implementation (adapter)
import (
	"golang-hexagonal-arch/module/book/domain/entity"

	"gorm.io/gorm"
)

type bookRepositoryGorm struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *bookRepositoryGorm {
	return &bookRepositoryGorm{db}
}

func (r *bookRepositoryGorm) GetListAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepositoryGorm) GetBookDetailId(id int) (entity.Book, error) {
	var book entity.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepositoryGorm) RegisterNewBook(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepositoryGorm) DeleteBook(id int) (entity.Book, error) {
	var book entity.Book
	err := r.db.Delete(&book, id).Error
	return book, err
}

func (r *bookRepositoryGorm) UpdateBook(id int) (entity.Book, error) {
	var book entity.Book

	err := r.db.First(&book, id).Error

	err = r.db.Save(&book).Error
	return book, err
}
