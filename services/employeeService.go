package services

import (
	"human-resource-api/models"
	"human-resource-api/repository"
)

type EmployeeService interface {
	FindAll() ([]models.EmployeeResponse, error)
	FindByID(employeeID int) (models.EmployeeResponse, error)
	Store(employeeRequest models.EmployeeRequest) (models.EmployeeResponse, error)
	Update(employeeRequest models.EmployeeRequest, employeeID int) (models.EmployeeResponse, error)
	Delete(employeeID int) (models.EmployeeResponse, error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) *employeeService {
	return &employeeService{employeeRepository}
}

func (s *employeeService) FindAll() ([]models.EmployeeResponse, error) {
	var employeesResponse []models.EmployeeResponse

	employees, err := s.employeeRepository.FindAll()
	if err != nil {
		return employeesResponse, err
	}

	for _, employee := range employees {
		newEmployeeResponse := models.EmployeeResponse{
			ID:           employee.ID,
			Name:         employee.Name,
			Age:          employee.Age,
			Sex:          employee.Sex,
			Phone:        employee.Phone,
			Email:        employee.Email,
			Address:      employee.Address,
			Picture:      employee.Picture,
			DepartmentID: employee.DepartmentID,
		}
		employeesResponse = append(employeesResponse, newEmployeeResponse)
	}

	return employeesResponse, err
}

func (s *employeeService) FindByID(employeeID int) (models.EmployeeResponse, error) {
	var employeeResponse models.EmployeeResponse

	employee, err := s.employeeRepository.FindByID(employeeID)
	if err != nil {
		return employeeResponse, err
	}

	employeeResponse.ID = employee.ID
	employeeResponse.Name = employee.Name
	employeeResponse.Age = employee.Age
	employeeResponse.Sex = employee.Sex
	employeeResponse.Phone = employee.Phone
	employeeResponse.Email = employee.Email
	employeeResponse.Address = employee.Address
	employeeResponse.Picture = employee.Picture
	employeeResponse.DepartmentID = employee.DepartmentID

	return employeeResponse, err
}

func (s *employeeService) Store(employeeRequest models.EmployeeRequest) (models.EmployeeResponse, error) {
	var employeeResponse models.EmployeeResponse

	newEmployee := models.Employee{
		Name:         employeeRequest.Name,
		Age:          employeeRequest.Age,
		Sex:          employeeRequest.Sex,
		Phone:        employeeRequest.Phone,
		Email:        employeeRequest.Email,
		Address:      employeeRequest.Address,
		Picture:      employeeRequest.Picture,
		DepartmentID: employeeRequest.DepartmentID,
	}

	storedEmployee, err := s.employeeRepository.Store(newEmployee)
	if err != nil {
		return employeeResponse, err
	}

	employeeResponse.ID = storedEmployee.ID
	employeeResponse.Name = storedEmployee.Name
	employeeResponse.Age = storedEmployee.Age
	employeeResponse.Sex = storedEmployee.Sex
	employeeResponse.Phone = storedEmployee.Phone
	employeeResponse.Email = storedEmployee.Email
	employeeResponse.Address = storedEmployee.Address
	employeeResponse.Picture = storedEmployee.Picture
	employeeResponse.DepartmentID = storedEmployee.DepartmentID

	return employeeResponse, err
}

func (s *employeeService) Update(employeeRequest models.EmployeeRequest, employeeID int) (models.EmployeeResponse, error) {
	var employeeResponse models.EmployeeResponse

	updateEmployee, err := s.employeeRepository.FindByID(employeeID)
	if err != nil {
		return employeeResponse, err
	}

	updateEmployee.Name = employeeRequest.Name
	updateEmployee.Age = employeeRequest.Age
	updateEmployee.Sex = employeeRequest.Sex
	updateEmployee.Phone = employeeRequest.Phone
	updateEmployee.Email = employeeRequest.Email
	updateEmployee.Address = employeeRequest.Address
	updateEmployee.Picture = employeeRequest.Picture
	updateEmployee.DepartmentID = employeeRequest.DepartmentID

	err = s.employeeRepository.Update(updateEmployee, employeeID)
	if err != nil {
		return employeeResponse, err
	}

	employeeResponse.ID = updateEmployee.ID
	employeeResponse.Name = updateEmployee.Name
	employeeResponse.Age = updateEmployee.Age
	employeeResponse.Sex = updateEmployee.Sex
	employeeResponse.Phone = updateEmployee.Phone
	employeeResponse.Email = updateEmployee.Email
	employeeResponse.Address = updateEmployee.Address
	employeeResponse.Picture = updateEmployee.Picture
	employeeResponse.DepartmentID = updateEmployee.DepartmentID

	return employeeResponse, err
}

func (s *employeeService) Delete(employeeID int) (models.EmployeeResponse, error) {
	var employeeResponse models.EmployeeResponse

	deleteEmployee, err := s.employeeRepository.FindByID(employeeID)
	if err != nil {
		return employeeResponse, err
	}

	err = s.employeeRepository.Delete(employeeID)
	if err != nil {
		return employeeResponse, err
	}

	employeeResponse.ID = deleteEmployee.ID
	employeeResponse.Name = deleteEmployee.Name
	employeeResponse.Age = deleteEmployee.Age
	employeeResponse.Sex = deleteEmployee.Sex
	employeeResponse.Phone = deleteEmployee.Phone
	employeeResponse.Email = deleteEmployee.Email
	employeeResponse.Address = deleteEmployee.Address
	employeeResponse.Picture = deleteEmployee.Picture
	employeeResponse.DepartmentID = deleteEmployee.DepartmentID

	return employeeResponse, err
}
