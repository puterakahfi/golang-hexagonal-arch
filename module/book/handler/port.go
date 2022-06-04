package handler

type BookHandlerInterface interface {
	GetListBooksHandler()
	GetDetailBookByIdHandler()
	DeleteBookHandler()
	RegisterRoutes()
}
