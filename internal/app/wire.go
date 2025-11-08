package app

import (
	"books-api/internal/adapters/db"
	"books-api/internal/adapters/https"
	"books-api/internal/core/book"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Init(dbConn *gorm.DB) *https.BookHandler {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	bookRepository := db.NewBookRepository(dbConn)
	bookService := book.NewBookService(bookRepository, logger)
	bookHandler := https.NewBookHandler(bookService, logger)
	return bookHandler
}
