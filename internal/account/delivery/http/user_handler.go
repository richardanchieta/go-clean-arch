package http

import (
	"context"
	"myapp/internal/account/usecase"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

type UserInput struct {
	Name     string `json:"name" doc:"Nome do usuário"`
	Email    string `json:"email" doc:"E-mail do usuário"`
	Password string `json:"password" doc:"Senha do usuário"`
}

type UserOutput struct {
	ID    string `json:"id" doc:"ID do usuário"`
	Name  string `json:"name" doc:"Nome do usuário"`
	Email string `json:"email" doc:"E-mail do usuário"`
}

func (h *UserHandler) RegisterUser(ctx context.Context, input *struct {
	Body UserInput
}) (*UserOutput, error) {
	user, err := h.usecase.RegisterUser(ctx, input.Body.Name, input.Body.Email, input.Body.Password)
	if err != nil {
		return nil, huma.Error500InternalServerError("Erro ao registrar usuário")
	}

	return &UserOutput{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
