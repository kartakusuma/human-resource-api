package routes

import (
	"human-resource-api/config"
	"human-resource-api/controllers"
	"human-resource-api/models"
	"human-resource-api/repository"
	"human-resource-api/services"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	db := config.Connection()
	models.MigrateModels(db)

	router := gin.Default()
	apiRoute := router.Group("/api")
	apiV1Route := apiRoute.Group("/v1")

	departmentRepository := repository.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepository)
	departmentController := controllers.NewDepartmentController(departmentService)

	apiV1Route.GET("/departments", departmentController.GetAll)
	apiV1Route.POST("/departments", departmentController.Store)
	apiV1Route.GET("/departments/:id", departmentController.GetByID)
	apiV1Route.PUT("/departments/:id", departmentController.Update)
	apiV1Route.DELETE("/departments/:id", departmentController.Destroy)

	return router
}
