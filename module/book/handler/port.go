package handler

type BookHandlerInterface interface {
	GetListBooksHandler()
	GetDetailBookByIdHandler()
	RegisterRoutes()
}
