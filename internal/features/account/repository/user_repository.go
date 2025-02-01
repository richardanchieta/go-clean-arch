package repository

import (
	"context"
	"myapp/internal/dal"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *dal.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: dal.New(pool),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, name, email, password string) (*dal.User, error) {
	user, err := r.queries.CreateUser(ctx, dal.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*dal.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
