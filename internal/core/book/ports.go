package book

type BookRepository interface {
	FindById(id int) (Book, error)
	FindAll() ([]Book, error)
}
type IBookService interface {
	GetBookByID(id int) (Book, error)
	GetAllBooks() ([]Book, error)
}
