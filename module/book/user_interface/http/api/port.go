package book

type BookHandlerInterface interface {
	GetListBooksHandler()
	GetDetailBookByIdHandler()
	DeleteBookHandler()
	RegisterRoutes()
	UpdateBookHandler()
}
