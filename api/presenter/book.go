package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
)

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Swagger error response
type ErrorResponse struct {
	Data   string `json:"data"`
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

// Swagger book response
type BookResponse struct {
	Data   Book `json:"data"`
	Status bool `json:"status"`
	Error  any  `json:"error"`
}

// Swagger books response
type BooksResponse struct {
	Data   []Book `json:"data"`
	Status bool   `json:"status"`
	Error  any    `json:"error"`
}

// Swagger book delete response
type DeleteResponse struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
	Error  any    `json:"error"`
}

func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func BooksSuccessResponse(books *[]entities.Book) *fiber.Map {
	var formattedBooks []Book
	for _, book := range *books {
		formattedBooks = append(formattedBooks, Book{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		})
	}

	return &fiber.Map{
		"status": true,
		"data":   formattedBooks,
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

func BookDeleteResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
		"error":  nil,
	}
}
