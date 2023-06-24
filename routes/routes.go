package routes

import "github.com/gin-gonic/gin"

func RouterSetup() *gin.Engine {
	router := gin.Default()

	apiRoute := router.Group("/api")

	apiV1Route := apiRoute.Group("/v1")
	apiV1Route.GET("/", func(ctx *gin.Context) {})

	return router
}
