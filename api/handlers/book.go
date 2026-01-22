package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vitaly06/shop-rest-api/api/presenter"
	"github.com/vitaly06/shop-rest-api/pkg/book"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
	"gorm.io/gorm"
)

// @Summary Get books
// @Description Find all books
// @Accept json
// @Produce json
// @Tags books
// @Success 200 {object} presenter.BooksResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /books [get]
func GetBooks(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchBooks()

		if err != nil {
			c.Status(500)
			return c.JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(presenter.BooksSuccessResponse(fetched))
	}
}

// @Summary Insert book
// @Accept json
// @Produce json
// @Tags books
// @Param data body presenter.CreateBookRequest true "Book data"
// @Success 201 {object} presenter.BookResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /books [post]
func InsertBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody presenter.CreateBookRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Title == "" || requestBody.Author == "" {
			return c.Status(400).JSON(presenter.BookErrorResponse(errors.New("Please specify title and author")))
		}

		book := &entities.Book{
			Title:  requestBody.Title,
			Author: requestBody.Author,
		}

		result, err := service.InsertBook(book)

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		return c.Status(201).JSON(presenter.BookSuccessResponse(result))
	}
}

// @Summary Get book
// @Description Find book by ID
// @Param id path int true "book ID"
// @Accept json
// @Produce json
// @Tags books
// @Success 200 {object} presenter.BookResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Router /books/{id} [get]
func GetBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookId, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}

		book, err := service.GetBook(int64(bookId))

		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(presenter.BookErrorResponse(errors.New("Book not found")))
		}

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		return c.Status(200).JSON(presenter.BookSuccessResponse(book))
	}
}

// @Summary Delete book
// @Description Delete book by ID
// @Param id path int true "book ID"
// @Accept json
// @Produce json
// @Tags books
// @Success 200 {object} presenter.DeleteResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookId, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}

		_, err = service.GetBook(int64(bookId))

		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(presenter.BookErrorResponse(errors.New("Book not found")))
		}

		err = service.DeleteBook(int64(bookId))

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		return c.Status(200).JSON(presenter.BookDeleteResponse())

	}
}

// @Summary Update book
// @Description Update book by ID
// @Param id path int true "book ID"
// @Param data body presenter.UpdateBookRequest true "Book data"
// @Accept json
// @Produce json
// @Tags books
// @Success 200 {object} presenter.BookResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /books/{id} [put]
func UpdateBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookId, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}

		_, err = service.GetBook(int64(bookId))

		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(presenter.BookErrorResponse(errors.New("Book not found")))
		}

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		var requestBody presenter.UpdateBookRequest
		err = c.BodyParser(&requestBody)

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}

		book := &entities.Book{
			ID:     uint(bookId),
			Title:  requestBody.Title,
			Author: requestBody.Author,
		}

		result, err := service.UpdateBook(book)

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		return c.JSON(presenter.BookSuccessResponse(result))
	}
}
