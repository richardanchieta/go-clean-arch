package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RegisterRoutes(api huma.API, router *gin.Engine) {
	huma.Register(api, huma.Operation{
		OperationID:   "registerUser",
		Summary:       "Registra um novo usuário",
		Method:        http.MethodPost,
		Path:          "/api/users",
		Tags:          []string{"account"},
		DefaultStatus: http.StatusCreated,
		Security: []map[string][]string{
			{"keycloak-openid": {"scope1"}},
		},
	}, h.RegisterUser)
}
