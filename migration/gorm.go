package migration

import (
	"golang-hexagonal-arch/module/book/domain/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Book{})
}
