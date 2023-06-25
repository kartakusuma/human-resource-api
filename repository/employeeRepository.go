package repository

import (
	"human-resource-api/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	FindByID(employeeID int) (models.Employee, error)
	Store(employee models.Employee) (models.Employee, error)
	Update(employee models.Employee, employeeID int) error
	Delete(employeeID int) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee

	err := r.db.Find(&employees).Error

	return employees, err
}

func (r *employeeRepository) FindByID(employeeID int) (models.Employee, error) {
	var employee models.Employee

	err := r.db.First(&employee, employeeID).Error

	return employee, err
}

func (r *employeeRepository) Store(employee models.Employee) (models.Employee, error) {
	err := r.db.Create(&employee).Error

	return employee, err
}

func (r *employeeRepository) Update(employee models.Employee, employeeID int) error {
	err := r.db.Model(&models.Employee{}).Where("ID = ?", employeeID).Updates(&employee).Error

	return err
}

func (r *employeeRepository) Delete(employeeID int) error {
	err := r.db.Delete(&models.Employee{}, employeeID).Error

	return err
}
