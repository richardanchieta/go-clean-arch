package repository

import (
	"context"
	"math/big"
	"myapp/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	queries *db.Queries
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		queries: db.New(pool),
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, userID uuid.UUID, productID uuid.UUID, quantity int, total float64) (*db.Order, error) {

	order, err := r.queries.CreateOrder(ctx, db.CreateOrderParams{
		UserID:    userID,
		ProductID: productID,
		Quantity:  int32(quantity),
		Total:     pgtype.Numeric{Int: big.NewInt(int64(total * 100)), Exp: 2, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &order, nil
}
