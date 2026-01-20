package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitaly06/shop-rest-api/api/handlers"
	"github.com/vitaly06/shop-rest-api/pkg/book"
)

func BookRouter(app *fiber.App, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.InsertBook(service))
}
