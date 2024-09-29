package modules

import (
	"blog-center/internal/handlers"
	"blog-center/internal/repository"
	"blog-center/internal/service"

	"go.uber.org/fx"
)

func RegisterUserDependencies() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.NewUserRepository,
			service.NewUserService,
			handlers.NewUserHandler,	
		),
	)
}