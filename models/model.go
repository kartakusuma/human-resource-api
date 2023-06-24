package models

import "gorm.io/gorm"

func MigrateModels(db *gorm.DB) {
	db.Debug().AutoMigrate(Department{})
	db.Debug().AutoMigrate(Employee{})
}
