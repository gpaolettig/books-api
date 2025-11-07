package http

import (
	"books-api/internal/core/book"
	"books-api/internal/core/shared"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookService book.IBookService
	logger      *zap.Logger
}

func NewBookHandler(bookService book.IBookService, logger *zap.Logger) *BookHandler {
	return &BookHandler{
		bookService: bookService,
		logger:      logger,
	}
}
func (h *BookHandler) handleGetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		h.logger.Error("invalid id", zap.Int("id", id))
		return
	}
	bk, err := h.bookService.GetBookByID(id)
	if err != nil {
		if errors.Is(err, shared.ErrBookNotFound) {
			h.logger.Warn("book not found", zap.Int("id", id))
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("failed to get book", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bk)

}
func (h *BookHandler) handleGetAllBooks(ctx *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		if errors.Is(err, shared.ErrNoBooksFound) {
			h.logger.Warn("no books found")
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("failed to get books", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}
