package book

import (
	"github.com/vitaly06/shop-rest-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	ReadBook() (*[]entities.Book, error)
	CreateBook(book *entities.Book) (*entities.Book, error)
	GetBook(id int64) (*entities.Book, error)
	DeleteBook(id int64) error
	UpdateBook(book *entities.Book) (*entities.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ReadBook() (*[]entities.Book, error) {
	var books []entities.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *repository) GetBook(id int64) (*entities.Book, error) {
	var book entities.Book

	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) DeleteBook(id int64) error {
	var book entities.Book

	if err := r.db.Delete(&book, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
