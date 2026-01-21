package book

import (
	"github.com/vitaly06/shop-rest-api/pkg/entities"
)

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]entities.Book, error)
	GetBook(id int64) (*entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchBooks() (*[]entities.Book, error) {
	return s.repository.ReadBook()
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}

func (s *service) GetBook(id int64) (*entities.Book, error) {
	return s.repository.GetBook(id)
}

func (s *service) DeleteBook(id int64) error {
	return s.repository.DeleteBook(id)
}

func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdateBook(book)
}
