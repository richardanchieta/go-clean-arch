package router

import (
	"net/http"

	"myapp/config"
	accountHttp "myapp/internal/features/account/delivery/http"
	checkoutHttp "myapp/internal/features/checkout/delivery/http"
	storefrontHttp "myapp/internal/features/storefront/delivery/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*http.Server
}

func SetupRouter(
	userHandler *accountHttp.UserHandler,
	productHandler *storefrontHttp.ProductHandler,
	orderHandler *checkoutHttp.OrderHandler,
	authConfig *config.AuthConfig,
) *Router {

	r := gin.Default()

	humaConfig := huma.DefaultConfig("My App API", "1.0.0")
	humaConfig.Info.Description = "# My App API"
	humaConfig.Info.TermsOfService = "https://app.myapp.com/terms"
	humaConfig.Info.Contact = &huma.Contact{
		Name:  "My App Contact",
		Email: "",
		URL:   "https://app.myapp.com/contact",
	}
	humaConfig.Servers = []*huma.Server{
		{
			URL:         "https://api.myapp.com",
			Description: "My App API for learning purposes",
		},
	}

	api := humagin.New(r, humaConfig)
	userHandler.RegisterRoutes(api, r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &Router{srv}
}
