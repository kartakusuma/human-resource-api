package main

import (
	"human-resource-api/config"
	"human-resource-api/models"
)

func main() {
	db := config.Connection()

	models.MigrateModels(db)
}
