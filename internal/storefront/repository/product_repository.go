package repository

import (
	"context"
	"math/big"
	"myapp/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	queries *db.Queries
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		queries: db.New(pool),
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, name, description string, price float64) (*db.Product, error) {
	product, err := r.queries.CreateProduct(ctx, db.CreateProductParams{
		Name:        name,
		Description: pgtype.Text{String: description, Valid: true},
		Price:       pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: 2, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) ListProducts(ctx context.Context) ([]db.Product, error) {
	products, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
