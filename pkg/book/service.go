package book

import (
	"github.com/vitaly06/shop-rest-api/api/presenter"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
)

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]presenter.Book, error)
	// UpdateBook(book *entities.Book) (*entities.Book, error)
	// RemoveBook(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchBooks() (*[]presenter.Book, error) {
	return s.repository.ReadBook()
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}
