package repository

import (
	"human-resource-api/models"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindAll() ([]models.Department, error)
	FindByID(departmentID int) (models.Department, error)
	Store(department models.Department) (models.Department, error)
	Update(department models.Department, departmentID int) error
	Delete(departmentID int) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department

	err := r.db.Find(&departments).Error

	return departments, err
}

func (r *departmentRepository) FindByID(departmentID int) (models.Department, error) {
	var department models.Department

	err := r.db.First(&department, departmentID).Error

	return department, err
}

func (r *departmentRepository) Store(department models.Department) (models.Department, error) {
	err := r.db.Create(&department).Error

	return department, err
}

func (r *departmentRepository) Update(department models.Department, departmentID int) error {
	err := r.db.Model(&models.Department{}).Where("ID = ?", departmentID).Updates(&department).Error

	return err
}

func (r *departmentRepository) Delete(departmentID int) error {
	err := r.db.Delete(&models.Department{}, departmentID).Error

	return err
}
