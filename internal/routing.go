package internal

import (
	"blog-center/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(
	r *gin.Engine, 
	userHandler *handlers.UserHandler,
	) {
		handlers.GroupUserHandlers(r, userHandler)
}
