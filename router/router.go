package router

import (
	"human-resource-api/config"
	"human-resource-api/controllers"
	"human-resource-api/models"
	"human-resource-api/repository"
	"human-resource-api/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterSetup() *gin.Engine {
	db := config.Connection()
	models.MigrateModels(db)

	router := gin.Default()

	// swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoute := router.Group("/api")
	apiV1Route := apiRoute.Group("/v1")

	// department routes
	departmentRepository := repository.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepository)
	departmentController := controllers.NewDepartmentController(departmentService)

	apiV1Route.GET("/departments", departmentController.GetAll)
	apiV1Route.POST("/departments", departmentController.Store)
	apiV1Route.GET("/departments/:id", departmentController.GetByID)
	apiV1Route.PUT("/departments/:id", departmentController.Update)
	apiV1Route.DELETE("/departments/:id", departmentController.Destroy)

	// employee routes
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	apiV1Route.GET("/employees", employeeController.GetAll)
	apiV1Route.POST("/employees", employeeController.Store)
	apiV1Route.GET("/employees/:id", employeeController.GetByID)
	apiV1Route.PUT("/employees/:id", employeeController.Update)
	apiV1Route.DELETE("/employees/:id", employeeController.Destroy)

	return router
}
