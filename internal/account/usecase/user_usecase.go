package usecase

import (
	"context"
	"errors"

	"myapp/db"
	"myapp/internal/account/repository"
)

// UserUsecase representa os casos de uso de usuário
type UserUsecase struct {
	userRepo *repository.UserRepository
}

// NewUserUsecase cria uma instância do caso de uso de usuário
func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// CreateUser lida com a criação de usuários
func (uc *UserUsecase) RegisterUser(ctx context.Context, name, email, password string) (*db.User, error) {
	existingUser, _ := uc.userRepo.GetUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("email já cadastrado")
	}
	user, err := uc.userRepo.CreateUser(ctx, name, email, password)

	return user, err
}
