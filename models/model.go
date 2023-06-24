package models

import "gorm.io/gorm"

func MigrateModels(db *gorm.DB) {
	db.Debug().AutoMigrate(Departemen{})
	db.Debug().AutoMigrate(Employee{})
}
