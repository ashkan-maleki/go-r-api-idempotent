package service

import (
	"context"
	"github.com/ashkan-maleki/go-r-api-idempotent/pkg/entity"
	"gorm.io/gorm"
)

type Shipping struct {
	db *gorm.DB
}

func NewShipping(db *gorm.DB) *Shipping {
	return &Shipping{db: db}
}

func (s *Shipping) Save(ctx context.Context, order entity.ShippingOrder) (entity.ShippingOrder, error) {
	err := s.db.WithContext(ctx).Save(&order).Error
	return order, err
}

func (s *Shipping) ByID(ctx context.Context, id uint) (*entity.ShippingOrder, error) {
	return s.by(ctx, "id", id)
}

func (s *Shipping) ByOrderID(ctx context.Context, id uint) (*entity.ShippingOrder, error) {
	return s.by(ctx, "order_id", id)
}

func (s *Shipping) by(ctx context.Context, key string, val any) (*entity.ShippingOrder, error) {
	var order entity.ShippingOrder
	if tr := s.db.WithContext(ctx).Where(key+"=?", val).First(&order); tr.Error != nil {
		return nil, tr.Error
	}
	return &order, nil
}
