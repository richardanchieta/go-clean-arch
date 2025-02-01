package usecase

import (
	"context"

	"myapp/internal/dal"
	"myapp/internal/features/checkout/repository"

	"github.com/google/uuid"
)

// OrderUsecase é um caso de uso de pedido
type OrderUsecase struct {
	orderRepo *repository.OrderRepository
}

// NewOrderUsecase cria um novo caso de uso de pedido
func NewOrderUsecase(orderRepo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo}
}

// CreateOrder cria um pedido
func (uc *OrderUsecase) CreateOrder(ctx context.Context, userID uuid.UUID, productID uuid.UUID, quantity int, total float64) (*dal.Order, error) {
	order, err := uc.orderRepo.CreateOrder(ctx, userID, productID, quantity, total)

	return order, err
}
