package book

import (
	"github.com/vitaly06/shop-rest-api/api/presenter"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	ReadBook() (*[]presenter.Book, error)
	CreateBook(book *entities.Book) (*entities.Book, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) ReadBook() (*[]presenter.Book, error) {
	var books []presenter.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	if err := r.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
