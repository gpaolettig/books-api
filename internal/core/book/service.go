package book

import "go.uber.org/zap"

type BookService struct {
	bookRepository BookRepository
	logger         *zap.Logger
}

func NewBookService(bookRepository BookRepository, logger *zap.Logger) *BookService {
	return &BookService{
		bookRepository: bookRepository,
		logger:         logger,
	}
}

func (s *BookService) GetBookByID(id int) (Book, error) {
	return s.bookRepository.FindById(id)
}

func (s *BookService) GetAllBooks() ([]Book, error) {
	return s.bookRepository.FindAll()
}
