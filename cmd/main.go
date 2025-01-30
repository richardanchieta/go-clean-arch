package main

import (
	"context"
	"log"
	"myapp/config"

	userHttp "myapp/internal/account/delivery/http"
	userRepo "myapp/internal/account/repository"
	userUC "myapp/internal/account/usecase"

	orderHttp "myapp/internal/checkout/delivery/http"
	orderRepo "myapp/internal/checkout/repository"
	orderUC "myapp/internal/checkout/usecase"

	productHttp "myapp/internal/storefront/delivery/http"
	productRepo "myapp/internal/storefront/repository"
	productUC "myapp/internal/storefront/usecase"

	"myapp/pkg/payments"
	"myapp/router"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	fx.Provide(
		config.LoadAuthConfig,
		config.NewRedisClient,
		config.NewDatabase,
		userRepo.NewUserRepository,
		userUC.NewUserUsecase,
		userHttp.NewUserHandler,
		productRepo.NewProductRepository,
		productUC.NewProductUsecase,
		productHttp.NewProductHandler,
		orderRepo.NewOrderRepository,
		orderUC.NewOrderUsecase,
		orderHttp.NewOrderHandler,
		payments.NewMockGateway,
		router.SetupRouter,
	),
	fx.Invoke(registerHooks),
)

func registerHooks(lc fx.Lifecycle, r *router.Router) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("🚀 Servidor rodando na porta 8080...")
				if err := r.ListenAndServe(); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("🛑 Servidor finalizado.")
			return nil
		},
	})
}

func main() {
	fx.New(AppModule).Run()
}
