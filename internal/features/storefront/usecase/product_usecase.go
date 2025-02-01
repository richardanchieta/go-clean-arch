package usecase

import (
	"context"

	"myapp/internal/dal"
	"myapp/internal/features/storefront/repository"
)

// UserUsecase representa os casos de uso de produto
type ProductUsecase struct {
	productRepo *repository.ProductRepository
}

// NewProductUsecase cria uma instância do caso de uso de produto
func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: productRepo}
}

// CreateProduct lida com a criação de produtos
func (uc *ProductUsecase) CreateProduct(ctx context.Context, name, description string, price float64) (*dal.Product, error) {
	product, err := uc.productRepo.CreateProduct(ctx, name, description, price)
	return product, err
}

// ListProducts lida com a listagem de produtos
func (uc *ProductUsecase) ListProducts(ctx context.Context) ([]dal.Product, error) {
	products, err := uc.productRepo.ListProducts(ctx)
	return products, err
}
