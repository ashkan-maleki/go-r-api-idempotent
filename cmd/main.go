package main

import (
	"github.com/ashkan-maleki/go-r-api-idempotent/cmd/config"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/db/pg"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/handler"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	db, err := pg.New(config.PostgresDsn)
	if err != nil {
		panic(err)
	}
	shippingRepo, err := service.NewShipping(db)
	if err != nil {
		panic(err)
	}
	shippingHandler := handler.NewShippingHandler(shippingRepo)

	app := fiber.New()
	app.Post("shipping/order", shippingHandler.Order)
	log.Fatal(app.Listen(":3000"))
}
