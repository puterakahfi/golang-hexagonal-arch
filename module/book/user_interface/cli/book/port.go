package book

// Book CLI/Command contract

type BookCommandInterface interface {
	GetListBooks()
	GetBookDetailById()
}
