package httpserver

func setRoutes() {
	//setAuthRoutes()
	setApiRoutes()
}

//func setAuthRoutes() {
//	//auth := routerServer.Group("/auth")
//
//	//auth.POST("/exchange", refreshTokenHandler)
//}

func setApiRoutes() {

	// status route
	routerServer.GET("/health", healthCheckHandler)

	//api := routerServer.Group("/api")

}
