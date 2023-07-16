package handler

import (
	"context"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/repo"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/service/idempotency"
	"github.com/ashkan-maleki/go-r-api-idempotent/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ShippingHandler struct {
	ShippingRepository  *repo.Shipping
	ShippingIdempotency *idempotency.Redis[entity.ShippingOrder]
}

func NewShippingHandler(shippingRepository *repo.Shipping, shippingIdempotency *idempotency.Redis[entity.ShippingOrder]) *ShippingHandler {
	return &ShippingHandler{ShippingRepository: shippingRepository, ShippingIdempotency: shippingIdempotency}
}

type PlaceShippingOrderRequest struct {
	OrderID string `json:"order_id"`
	Vendor  string `json:"vendor"`
	Address string `json:"address"`
}

func (s *ShippingHandler) Order(c *fiber.Ctx) error {
	var request PlaceShippingOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	shippingOrder, exists, err := s.ShippingIdempotency.Start(context.Background(), request.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if exists {
		return c.Status(fiber.StatusOK).JSON(map[string]any{
			"ok":          true,
			"shipping_id": shippingOrder.ID,
		})
	}

	<-time.After(time.Second * 3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createdOrder, err := s.ShippingRepository.Save(ctx, entity.ShippingOrder{
		OrderID: request.OrderID,
		Vendor:  request.Vendor,
		Address: request.Address,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := s.ShippingIdempotency.Store(context.Background(), createdOrder.OrderID, createdOrder); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"ok":          true,
		"shipping_id": createdOrder.ID,
	})

}
