package api

// Handler API port
type BookHandlerInterface interface {
	GetListBooksHandler()
	GetDetailBookByIdHandler()
	DeleteBookHandler()
	RegisterRoutes()
	UpdateBookHandler()
}
