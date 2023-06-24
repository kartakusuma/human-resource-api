package services

import (
	"human-resource-api/models"
	"human-resource-api/repository"
)

type DepartmentService interface {
	FindAll() ([]models.DepartmentResponse, error)
	FindByID(departmentID int) (models.DepartmentResponse, error)
	Store(department models.DepartmentRequest) (models.DepartmentResponse, error)
	Update(department models.DepartmentRequest) (models.DepartmentResponse, error)
	Delete(departmentID int) (models.DepartmentResponse, error)
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRepository repository.DepartmentRepository) *departmentService {
	return &departmentService{departmentRepository}
}

func (s *departmentService) FindAll() ([]models.DepartmentResponse, error) {
	var departmentsResponse []models.DepartmentResponse

	departments, err := s.departmentRepository.FindAll()
	if err != nil {
		return departmentsResponse, err
	}

	for _, department := range departments {
		newDepartmentResponse := models.DepartmentResponse{
			ID:   department.ID,
			Name: department.Name,
			City: department.City,
		}

		departmentsResponse = append(departmentsResponse, newDepartmentResponse)
	}

	return departmentsResponse, err
}

func (s *departmentService) FindByID(departmentID int) (models.DepartmentResponse, error) {
	var departmentResponse models.DepartmentResponse

	department, err := s.departmentRepository.FindByID(departmentID)
	if err != nil {
		return departmentResponse, err
	}

	departmentResponse.ID = department.ID
	departmentResponse.Name = department.Name
	departmentResponse.City = department.City

	return departmentResponse, err
}

func (s *departmentService) Store(departmentRequest models.DepartmentRequest) (models.DepartmentResponse, error) {
	var departmentResponse models.DepartmentResponse

	newDepartment := models.Department{
		Name: departmentRequest.Name,
		City: departmentRequest.City,
	}

	storedDepartment, err := s.departmentRepository.Store(newDepartment)
	if err != nil {
		return departmentResponse, err
	}

	departmentResponse.ID = storedDepartment.ID
	departmentResponse.Name = storedDepartment.City
	departmentResponse.City = storedDepartment.City

	return departmentResponse, err
}

func (s *departmentService) Update(departmentRequest models.DepartmentRequest) (models.DepartmentResponse, error) {
	var departmentResponse models.DepartmentResponse

	newDepartment := models.Department{
		Name: departmentRequest.Name,
		City: departmentRequest.City,
	}

	updatedDepartment, err := s.departmentRepository.Store(newDepartment)
	if err != nil {
		return departmentResponse, err
	}

	departmentResponse.ID = updatedDepartment.ID
	departmentResponse.Name = updatedDepartment.Name
	departmentResponse.City = updatedDepartment.City

	return departmentResponse, err
}

func (s *departmentService) Delete(departmentID int) (models.DepartmentResponse, error) {
	var departmentResponse models.DepartmentResponse

	department, err := s.departmentRepository.FindByID(departmentID)
	if err != nil {
		return departmentResponse, err
	}

	deletedDepartment, err := s.departmentRepository.Delete(department)
	if err != nil {
		return departmentResponse, err
	}

	departmentResponse.ID = deletedDepartment.ID
	departmentResponse.Name = deletedDepartment.Name
	departmentResponse.City = deletedDepartment.City

	return departmentResponse, err
}
