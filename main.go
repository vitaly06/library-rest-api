package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vitaly06/shop-rest-api/api/routes"
	"github.com/vitaly06/shop-rest-api/internal/infrastructure/config"
	"github.com/vitaly06/shop-rest-api/internal/infrastructure/database"
	"github.com/vitaly06/shop-rest-api/pkg/book"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg)
	// Book init
	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	if err != nil {
		log.Fatal("Database Connection Error: " + err.Error())
	}

	database.AutoMigrate(db)

	app := fiber.New()
	app.Use(cors.New())

	routes.BookRouter(app, bookService)

	fmt.Printf("App successfully started on port: %s\n", cfg.Port)
	app.Listen(":" + cfg.Port)

}
