package usecase

import (
	"context"
	"errors"
	"fmt"
)

type RegisterUserInput struct {
	Name     string `json:"name" doc:"Nome do usuário"`
	Email    string `json:"email" doc:"E-mail do usuário"`
	Password string `json:"password" doc:"Senha do usuário"`
}

type RegisterUserOutput struct {
	ID    string `json:"id" doc:"ID do usuário"`
	Name  string `json:"name" doc:"Nome do usuário"`
	Email string `json:"email" doc:"E-mail do usuário"`
}

// CreateUser lida com a criação de usuários
func (uc *UserUsecase) RegisterUser(ctx context.Context, input RegisterUserInput) (*RegisterUserOutput, error) {
	existingUser, _ := uc.userRepo.GetUserByEmail(ctx, input.Email)
	if existingUser != nil {
		return nil, errors.New("email já cadastrado")
	}

	user, err := uc.userRepo.CreateUser(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		err = fmt.Errorf("erro ao criar usuário %w", err)
		return nil, err
	}

	output := &RegisterUserOutput{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}

	return output, err
}
