package http

import (
	"context"
	"myapp/internal/storefront/usecase"

	"github.com/danielgtaylor/huma/v2"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

type ListProductsInput struct {
}

type ListProductsOutput struct {
	Items []ProductDto `json:"items" doc:"Lista de produtos"`
}

type ProductDto struct {
	ID          string  `json:"id" doc:"ID do produto"`
	Name        string  `json:"name" doc:"Nome do produto"`
	Description string  `json:"description" doc:"Descrição do produto"`
	Price       float64 `json:"price" doc:"Preço do produto"`
}

func (h *ProductHandler) ListProducts(ctx context.Context, input struct {
	Body ListProductsInput
}) (*ListProductsOutput, error) {
	products, err := h.usecase.ListProducts(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("Erro ao buscar produtos")
	}

	output := &ListProductsOutput{}
	for _, product := range products {

		price, _ := product.Price.Int.Float64()

		output.Items = append(output.Items, ProductDto{
			ID:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description.String,
			Price:       price,
		})
	}

	return output, nil
}
