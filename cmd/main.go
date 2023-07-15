package main

import (
	"github.com/ashkan-maleki/go-r-api-idempotent/cmd/config"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/db/pg"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/db/pg/redis"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/handler"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/service"
	"github.com/ashkan-maleki/go-r-api-idempotent/pkg/entity"
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

	shippingIdempotency := redis.NewRedis[entity.ShippingOrder](config.RedisUrl)

	shippingHandler := handler.NewShippingHandler(shippingRepo, shippingIdempotency)

	app := fiber.New()
	app.Post("shipping/order", shippingHandler.Order)
	log.Fatal(app.Listen(":3000"))
}
