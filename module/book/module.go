package book

import (
	"golang-hexagonal-arch/module/book/repository"
	"golang-hexagonal-arch/module/book/service"
	book_api "golang-hexagonal-arch/module/book/user_interface/http/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookModule struct {
	service    service.BookServiceInterface
	handler    book_api.BookHandlerInterface
	repository repository.BookRepositoryInterface
}

func InitModule(db *gorm.DB, gin *gin.Engine) *bookModule {
	bookModule := &bookModule{}
	bookModule.repository = repository.NewGormRepository(db)
	bookModule.service = service.NewService(bookModule.repository)
	bookModule.handler = book_api.NewGinHandler(gin, bookModule.service)
	bookModule.handler.RegisterRoutes()
	return bookModule
}
