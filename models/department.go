package models

import "time"

type Department struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(50)"`
	City      string     `json:"city" gorm:"type:varchar(50)"`
	Employees []Employee `json:"employees"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type DepartmentRequest struct {
	Name string `json:"name" gorm:"type:varchar(50)"`
	City string `json:"city" gorm:"type:varchar(50)"`
}

type DepartmentResponse struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50)"`
	City string `json:"city" gorm:"type:varchar(50)"`
}
