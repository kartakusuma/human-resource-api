package controllers

import (
	"human-resource-api/models"
	"human-resource-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type employeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) *employeeController {
	return &employeeController{employeeService}
}

// GetAllEmployees	godoc
// @Summary	Get all employees
// @Description Get all employees
// @Produce application/json
// @Tags employee
// @Success 200 {object} models.EmployeeResponse{}
// @Router /employees [get]
func (c *employeeController) GetAll(ctx *gin.Context) {
	employees, err := c.employeeService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

// GetEmployeeByID	godoc
// @Summary	Get an employee
// @Description Get an employee by ID
// @Param employeeID path int true "Get employee by ID"
// @Produce application/json
// @Tags employee
// @Success 200 {object} models.EmployeeResponse{}
// @Router /employees/{employeeID} [get]
func (c *employeeController) GetByID(ctx *gin.Context) {
	employeeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee, err := c.employeeService.FindByID(employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

// StoreEmployee	godoc
// @Summary	Store employee
// @Description Store a new employee in database
// @Param employee body models.EmployeeRequest true "Store employee"
// @Produce application/json
// @Tags employee
// @Success 200 {object} models.EmployeeResponse{}
// @Router /employees [post]
func (c *employeeController) Store(ctx *gin.Context) {
	var employeeRequest models.EmployeeRequest

	err := ctx.ShouldBindJSON(&employeeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	storedEmployee, err := c.employeeService.Store(employeeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": storedEmployee,
	})
}

// UpdateEmployee	godoc
// @Summary	Update employee
// @Description Update a employee in database
// @Param employeeID path int true "Update employee by ID"
// @Param employee body models.EmployeeRequest true "Update employee by ID"
// @Produce application/json
// @Tags employee
// @Success 200 {object} models.EmployeeResponse{}
// @Router /employees/{employeeID} [put]
func (c *employeeController) Update(ctx *gin.Context) {
	var employeeRequest models.EmployeeRequest

	employeeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&employeeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedEmployee, err := c.employeeService.Update(employeeRequest, employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedEmployee,
	})
}

// DestroyEmployee	godoc
// @Summary	Destroy employee
// @Description Delete a employee in database
// @Param employeeID path int true "Delete employee by ID"
// @Produce application/json
// @Tags employee
// @Success 200 {object} models.EmployeeResponse{}
// @Router /employees/{employeeID} [delete]
func (c *employeeController) Destroy(ctx *gin.Context) {
	employeeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletedEmployee, err := c.employeeService.Delete(employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": deletedEmployee,
	})
}
