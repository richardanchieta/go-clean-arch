package usecase

import (
	"myapp/internal/features/account/repository"
)

// UserUsecase representa os casos de uso de usuário
type UserUsecase struct {
	userRepo *repository.UserRepository
}

// NewUserUsecase cria uma instância do caso de uso de usuário
func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}
