package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"mado/internal/config"
	"mado/internal/controller/http/middleware"
	"mado/internal/core"
)

// Deps is a http handler dependencies.
type Deps struct {
	Logger   *zap.Logger
	Services core.Services
}

// NewRouter returns a new http router.
func NewRouter(deps Deps) *gin.Engine {
	router := gin.New()

	if config.Get().IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	middleware.ApplyMiddlewares(router, deps.Logger)

	api := router.Group("/api")
	{

		newUserHandler(userDeps{
			router: api,

			userService: deps.Services.User,
		})

		// newProfileHandler(profileDeps{
		// 	router:         api,

		// 	profileService: deps.Services.Profile,
		// })
	}

	return router
}
