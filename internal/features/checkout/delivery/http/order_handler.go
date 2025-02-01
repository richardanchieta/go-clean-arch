package http

import (
	"context"

	"myapp/internal/features/checkout/usecase"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(u *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: u}
}

type CreateOrderInput struct {
	UserID    string  `json:"user_id" doc:"ID do usuário"`
	ProductID string  `json:"product_id" doc:"ID do produto"`
	Quantity  int     `json:"quantity" doc:"Quantidade"`
	Total     float64 `json:"total" doc:"Total"`
}

type CreateOrderOutput struct {
	ID        string  `json:"id" doc:"ID do pedido"`
	UserID    string  `json:"user_id" doc:"ID do usuário"`
	ProductID string  `json:"product_id" doc:"ID do produto"`
	Quantity  int     `json:"quantity" doc:"Quantidade"`
	Total     float64 `json:"total" doc:"Total"`
}

func (h *OrderHandler) CreateOrder(ctx context.Context, input struct {
	Body CreateOrderInput
}) (*CreateOrderOutput, error) {

	userUUID := uuid.MustParse(input.Body.UserID)
	productUUID := uuid.MustParse(input.Body.ProductID)

	order, err := h.usecase.CreateOrder(
		ctx,
		userUUID,
		productUUID,
		input.Body.Quantity,
		input.Body.Total,
	)
	if err != nil {
		return nil, huma.Error500InternalServerError("Erro ao processar pedido")
	}

	total, _ := order.Total.Int.Float64()
	output := &CreateOrderOutput{
		ID:        order.ID.String(),
		UserID:    order.UserID.String(),
		ProductID: order.ProductID.String(),
		Quantity:  int(order.Quantity),
		Total:     float64(total / 100),
	}

	return output, nil
}
