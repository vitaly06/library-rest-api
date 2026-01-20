package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func BooksSuccessResponse(books *[]Book) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   books,
		"error":  nil,
	}
}

func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book := Book{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
	}

	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}
