package tests

import (
	"books-api/internal/adapters/https"
	"books-api/internal/app"
	"books-api/internal/core/book"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckGetBookAndGetAllBooks(t *testing.T) {
	r := setup()
	t.Run("HealthcheckResponseOk", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)
		res := performRequest(r, req)
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{"message":"UP"}`, res.Body.String())
	})
	t.Run("TestGetBookByIDResponseOK", func(t *testing.T) {
		var resBook *book.Book
		req, _ := http.NewRequest(http.MethodGet, "/books/1", nil)
		res := performRequest(r, req)
		require.NotNil(t, res)
		require.Equal(t, http.StatusOK, res.Code)
		require.NoError(t, json.Unmarshal(res.Body.Bytes(), &resBook))
		require.Equal(t, 1, resBook.ID)
		require.Equal(t, "Clean Code", resBook.Title)
		require.Equal(t, "Robert C. Martin", resBook.Author)
		require.Equal(t, "978-0132350884", resBook.ISBN)
	})
	t.Run("TestGetBookByIDResponseNotFound", func(t *testing.T) {
		//TODO
	})
	t.Run("TestGetAllBooksResponseOK", func(t *testing.T) {
		var resBooks []*book.Book
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		res := performRequest(r, req)
		require.NotNil(t, res)
		require.Equal(t, http.StatusOK, res.Code)
		require.NoError(t, json.Unmarshal(res.Body.Bytes(), &resBooks))
		require.Equal(t, 1, resBooks[0].ID)
		require.Equal(t, "Clean Code", resBooks[0].Title)
		require.Equal(t, "Robert C. Martin", resBooks[0].Author)
		require.Equal(t, "978-0132350884", resBooks[0].ISBN)
		require.Equal(t, 2, resBooks[1].ID)
		require.Equal(t, "The Pragmatic Programmer", resBooks[1].Title)
		require.Equal(t, "Andrew Hunt", resBooks[1].Author)
		require.Equal(t, "978-0201616224", resBooks[1].ISBN)
	})
	t.Run("TestGetAllBooksResponseNotFound", func(t *testing.T) {
		//TODO
	})

}
func setup() *gin.Engine {
	dbConn, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err := dbConn.AutoMigrate(&book.Book{}); err != nil {
		panic(fmt.Errorf("Failed to migrate db: %v", err))
	}
	dbConn.Create(&book.Book{ID: 1, Title: "Clean Code", Author: "Robert C. Martin", ISBN: "978-0132350884"})
	dbConn.Create(&book.Book{ID: 2, Title: "The Pragmatic Programmer", Author: "Andrew Hunt", ISBN: "978-0201616224"})
	bookHandler := app.Init(dbConn)
	r := gin.Default()
	https.InitRoutes(r, bookHandler)
	return r
}

func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
