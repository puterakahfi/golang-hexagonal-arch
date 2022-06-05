package book

import (
	"golang-hexagonal-arch/module/book/domain"
	"golang-hexagonal-arch/module/book/domain/usecase"
	"golang-hexagonal-arch/module/book/repository"
	book_api "golang-hexagonal-arch/module/book/user_interface/http/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookModule struct {
	usecase    usecase.BookUseCaseInterface
	handler    book_api.BookHandlerInterface
	repository domain.BookRepositoryInterface
}

func InitModule(db *gorm.DB, gin *gin.Engine) *bookModule {
	bookModule := &bookModule{}
	bookModule.repository = repository.NewGormRepository(db)
	bookModule.usecase = usecase.NewBookUseCase(bookModule.repository)
	bookModule.handler = book_api.NewGinHandler(gin, bookModule.usecase)
	bookModule.handler.RegisterRoutes()
	return bookModule
}
