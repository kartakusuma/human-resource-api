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
