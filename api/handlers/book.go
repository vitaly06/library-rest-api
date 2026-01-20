package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vitaly06/shop-rest-api/api/presenter"
	"github.com/vitaly06/shop-rest-api/pkg/book"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
)

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

func InsertBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)

		if err != nil {
			return c.Status(400).JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Title == "" || requestBody.Author == "" {
			return c.Status(400).JSON(presenter.BookErrorResponse(errors.New("Please specify title and author")))
		}

		result, err := service.InsertBook(&requestBody)

		if err != nil {
			return c.Status(500).JSON(presenter.BookErrorResponse(err))
		}

		return c.Status(200).JSON(presenter.BookSuccessResponse(result))
	}
}
