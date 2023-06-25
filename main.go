package main

import (
	_ "human-resource-api/docs"
	"human-resource-api/router"
)

// @title Human Resource API
// @version 1.0
// @description A REST API for Human Resource Information System using Gin, GORM, and MySQL

// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := router.RouterSetup()

	router.Run(":8080")
}
