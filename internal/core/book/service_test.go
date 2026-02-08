package book

import (
	"books-api/internal/core/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_GetBookByID(t *testing.T) {
	type fields struct {
		bookRepository BookRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    func(t *testing.T, book Book)
		wantErr func(t *testing.T, err error)
	}{
		{
			name: "error book not found",
			fields: fields{
				bookRepository: &mockBookRepository{
					mockFindById: func(id int) (Book, error) {
						return Book{}, shared.ErrBookNotFound
					},
				},
			},
			args: args{
				id: 1,
			},
			want: nil,
			wantErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
				require.EqualError(t, err, shared.ErrBookNotFound.Error())
			},
		},
		{
			name: "success book found",
			fields: fields{
				bookRepository: &mockBookRepository{
					mockFindById: func(id int) (Book, error) {
						return Book{
							ID:     1,
							Title:  "The Go Programming Language",
							Author: "Alan A. A. Donovan",
							ISBN:   "978-0134190440",
						}, nil
					},
				},
			},
			args: args{
				id: 1,
			},
			want: func(t *testing.T, book Book) {
				require.NotNil(t, book)
				require.Equal(t, 1, book.ID)
				require.Equal(t, "The Go Programming Language", book.Title)
				require.Equal(t, "Alan A. A. Donovan", book.Author)
				require.Equal(t, "978-0134190440", book.ISBN)
			},
			wantErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BookService{
				bookRepository: tt.fields.bookRepository,
			}
			got, err := s.GetBookByID(tt.args.id)

			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			if tt.want != nil {
				tt.want(t, got)
			}
		})
	}
}

type mockBookRepository struct {
	mockFindById func(id int) (Book, error)
	mockFindAll  func() ([]Book, error)
}

// implement the BookRepository interface
func (m *mockBookRepository) FindById(id int) (Book, error) {
	return m.mockFindById(id)
}
func (m *mockBookRepository) FindAll() ([]Book, error) {
	return m.mockFindAll()
}
