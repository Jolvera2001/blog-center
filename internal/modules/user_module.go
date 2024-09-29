package modules

import (
	"blog-center/internal/domain"
	"blog-center/internal/handlers"
	"blog-center/internal/repository"
	"blog-center/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterUserDependencies() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				repository.NewUserRepository,
				fx.As(new(domain.IUserRepository)),
			),
			fx.Annotate(
				service.NewUserService,
				fx.As(new(service.IUserService)),
			),
			handlers.NewUserHandler,
		),
		fx.Invoke(func(engine *gin.Engine, handler *handlers.UserHandler) {
			handlers.GroupUserHandlers(engine, handler)
		}),
	)
}
