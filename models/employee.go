package models

import "time"

type Employee struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"type:varchar(50)"`
	Age          int       `json:"age" gorm:"type:integer(2)"`
	Sex          string    `json:"sex" gorm:"type:varchar(20)"`
	Phone        int       `json:"phone" gorm:"type:integer(15)"`
	Email        string    `json:"email" gorm:"type:varchar(50)"`
	Address      string    `json:"address" gorm:"type:text"`
	Picture      string    `json:"picture" gorm:"type:varchar(255)"`
	DepartmentID int       `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	Name         string `json:"name" gorm:"type:varchar(50)"`
	Age          int    `json:"age" gorm:"type:integer(2)"`
	Sex          string `json:"sex" gorm:"type:varchar(20)"`
	Phone        int    `json:"phone" gorm:"type:integer(15)"`
	Email        string `json:"email" gorm:"type:varchar(50)"`
	Address      string `json:"address" gorm:"type:text"`
	Picture      string `json:"picture" gorm:"type:varchar(255)"`
	DepartmentID int    `json:"department_id"`
}

type EmployeeResponse struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"type:varchar(50)"`
	Age          int    `json:"age" gorm:"type:integer(2)"`
	Sex          string `json:"sex" gorm:"type:varchar(20)"`
	Phone        int    `json:"phone" gorm:"type:integer(15)"`
	Email        string `json:"email" gorm:"type:varchar(50)"`
	Address      string `json:"address" gorm:"type:text"`
	Picture      string `json:"picture" gorm:"type:varchar(255)"`
	DepartmentID int    `json:"department_id"`
}
