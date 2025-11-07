package app

import (
	"books-api/internal/adapters/db"
	"books-api/internal/adapters/http"
	"books-api/internal/core/book"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Init(dbConn *gorm.DB) *http.BookHandler {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	bookRepository := db.NewBookRepository(dbConn)
	bookService := book.NewBookService(bookRepository, logger)
	bookHandler := http.NewBookHandler(bookService, logger)
	return bookHandler
}
