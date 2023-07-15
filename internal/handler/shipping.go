package handler

import (
	"context"
	"github.com/ashkan-maleki/go-r-api-idempotent/internal/service"
	"github.com/ashkan-maleki/go-r-api-idempotent/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ShippingHandler struct {
	ShippingRepository service.Shipping
}

func NewShippingHandler(shippingRepository service.Shipping) *ShippingHandler {
	return &ShippingHandler{ShippingRepository: shippingRepository}
}

type PlaceShippingOrderRequest struct {
	OrderID string `json:"order_id"`
	Vendor  string `json:"vendor"`
	Address string `json:"address"`
}

func (s *ShippingHandler) Order(c *fiber.Ctx) {
	<-time.After(time.Second * 2)
	var request PlaceShippingOrderRequest
	if err := c.BodyParser(&request); err != nil {
		_ = c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// ====

	createdOrder, err := s.ShippingRepository.Save(ctx, entity.ShippingOrder{
		OrderID: request.OrderID,
		Vendor:  request.Vendor,
		Address: request.Address,
	})
	if err != nil {
		_ = c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	_ = c.Status(fiber.StatusCreated).JSON(map[string]any{
		"ok":          true,
		"shipping_id": createdOrder.ID,
	})

}
