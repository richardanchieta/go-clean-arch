package repository

import (
	"context"
	"math/big"
	"myapp/internal/dal"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	queries *dal.Queries
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		queries: dal.New(pool),
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, name, description string, price float64) (*dal.Product, error) {
	product, err := r.queries.CreateProduct(ctx, dal.CreateProductParams{
		Name:        name,
		Description: pgtype.Text{String: description, Valid: true},
		Price:       pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: 2, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) ListProducts(ctx context.Context) ([]dal.Product, error) {
	products, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
