package controllers

import (
	"human-resource-api/models"
	"human-resource-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type departmentController struct {
	departmentService services.DepartmentService
}

func NewDepartmentController(departmentService services.DepartmentService) *departmentController {
	return &departmentController{departmentService}
}

func (c *departmentController) GetAll(ctx *gin.Context) {
	departments, err := c.departmentService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": departments,
	})
}

func (c *departmentController) GetByID(ctx *gin.Context) {
	departmentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	department, err := c.departmentService.FindByID(departmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": department,
	})
}

func (c *departmentController) Store(ctx *gin.Context) {
	var departmentRequest models.DepartmentRequest

	err := ctx.ShouldBindJSON(&departmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	storedDepartment, err := c.departmentService.Store(departmentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": storedDepartment,
	})
}

func (c *departmentController) Update(ctx *gin.Context) {
	var departmentRequest models.DepartmentRequest

	departmentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&departmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedDepartment, err := c.departmentService.Update(departmentRequest, departmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedDepartment,
	})
}

func (c departmentController) Destroy(ctx *gin.Context) {
	departmentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletedDepartment, err := c.departmentService.Delete(departmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": deletedDepartment,
	})
}
