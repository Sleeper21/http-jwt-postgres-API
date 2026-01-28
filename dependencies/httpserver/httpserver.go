package httpserver

import (
	"core/app/domain/services"
	"core/app/domain/services/user"
	"fmt"

	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func CreateHttpServer(logger services.Logger, userRepo user.UserRepository) (*gin.Engine, string, error) {
	if os.Getenv("DD_ENV") != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	serverConfigs, err := LoadServerConfigs()
	if err != nil {
		logger.WithError(err, "Error loading server configs")
		return nil, "", err
	}
	logger.Infof("Server configs loaded: %s", serverConfigs.ApiHost)

	router := gin.New()
	//router.SetTrustedProxies(nil)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	//router.Use(loggerMiddleware(logger))

	// Create and set up the router
	httpRouter := CreateHttpRouter(logger, userRepo)
	httpRouter.SetRoutes(router)

	return router, serverConfigs.ApiHost, nil
}

func Run(router *gin.Engine, addr string) error {
	err := router.Run(addr)
	if err != nil {
		return fmt.Errorf("failed to start router: %w", err)
	}
	return nil
}
