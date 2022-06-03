package book

import (
	"golang-hexagonal-arch/module/book/handler"
	"golang-hexagonal-arch/module/book/repository"
	"golang-hexagonal-arch/module/book/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookModule struct {
	service    service.BookServiceInterface
	handler    handler.BookHandlerInterface
	repository repository.BookRepositoryInterface
}

func InitModule(db *gorm.DB, gin *gin.Engine) *bookModule {
	bookModule := &bookModule{}
	bookModule.repository = repository.NewGormRepository(db)
	bookModule.service = service.NewService(bookModule.repository)
	bookModule.handler = handler.NewGinHandler(gin, bookModule.service)
	bookModule.handler.RegisterRoutes()
	return bookModule
}
