package config

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthConfig armazena as configurações do Keycloak
type AuthConfig struct {
	KeycloakURL   string
	Realm         string
	ClientID      string
	ClientSecret  string
	PublicKey     string
	GoCloakClient *gocloak.GoCloak
	Context       context.Context
}

// LoadAuthConfig inicializa autenticação com Keycloak
func LoadAuthConfig() *AuthConfig {

	client := gocloak.NewClient("https://id.myapp.com")

	auth := &AuthConfig{
		KeycloakURL:   "https://id.myapp.com",
		Realm:         "myapp",
		ClientID:      "myapp-app",
		ClientSecret:  "myapp-app-secret",
		GoCloakClient: client,
		Context:       context.Background(),
	}

	// Obtendo chave pública
	pubKey, err := auth.GoCloakClient.GetCerts(auth.Context, auth.Realm)
	if err != nil {
		log.Fatalf("Erro ao buscar chave pública do Keycloak: %v", err)
	}
	auth.PublicKey = (*(*pubKey.Keys)[0].X5c)[0]

	return auth
}

// JWTMiddleware para proteção de rotas
func (auth *AuthConfig) JWTMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		tokenString := ctx.Header("Authorization")
		if tokenString == "" {
			hErr := huma.Error401Unauthorized("Token ausente")
			_ = huma.WriteErr(api, ctx, hErr.GetStatus(), hErr.Error())
			return
		}
		tokenString = tokenString[len("Bearer "):]
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.PublicKey), nil
		})
		if err != nil {
			hErr := huma.Error401Unauthorized("Token inválido")
			_ = huma.WriteErr(api, ctx, hErr.GetStatus(), hErr.Error())
			return
		}
		next(ctx)
	}
}
