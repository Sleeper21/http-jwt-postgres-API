package httpserver

import (
	"core/app/domain"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var routerServer *gin.Engine
var serverConfigs ServerConfigs

//var serverDbClient domain.DbClient

func CreateHttpServer(logger domain.Logger) (string, error) {
	if os.Getenv("DD_ENV") != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	serverConfigs, err := LoadServerConfigs()
	if err != nil {
		logger.WithError(err, "Error loading server configs")
		return "", err
	}
	logger.Infof("Server configs loaded: %s", serverConfigs.ApiHost)

	routerServer = gin.New()

	// Enable CORS for VUE frontend
	// routerServer.Use(cors.New(cors.Config{
	// 	//AllowOrigins:     []string{serverConfigs.VueSPA}, // This is the frontend app URL
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	routerServer.SetTrustedProxies(nil)

	//routerServer.Use(loggerMiddleware(logger))

	//serverDbClient = dbClient

	setRoutes()

	return serverConfigs.ApiHost, nil
}

func Run(addr string) error {

	err := routerServer.Run(addr)
	if err != nil {
		return fmt.Errorf("failed to start router: %w", err)
	}

	fmt.Println("Server started on: ", serverConfigs.ApiHost)

	return nil
}
