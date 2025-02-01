package http

import (
	"context"
	"fmt"
	"myapp/internal/features/account/usecase"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

type BodyWrapper[T any] struct {
	Body T
}

func (h *UserHandler) RegisterUser(
	ctx context.Context,
	input *BodyWrapper[usecase.RegisterUserInput],
) (*BodyWrapper[*usecase.RegisterUserOutput], error) {
	user, err := h.usecase.RegisterUser(ctx, input.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, huma.Error500InternalServerError("Erro ao registrar usuário")
	}

	output := &BodyWrapper[*usecase.RegisterUserOutput]{Body: user}

	return output, nil
}
