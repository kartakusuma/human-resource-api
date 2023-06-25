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

// GetAllDepartments	godoc
// @Summary	Get all departments
// @Description Get all departments
// @Produce application/json
// @Tags department
// @Success 200 {object} models.DepartmentResponse{}
// @Router /departments [get]
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

// GetDepartmentByID	godoc
// @Summary	Get a department
// @Description Get a department by ID
// @Param departmentID path int true "Get department by ID"
// @Produce application/json
// @Tags department
// @Success 200 {object} models.DepartmentResponse{}
// @Router /departments/{departmentID} [get]
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

// StoreDepartment	godoc
// @Summary	Store department
// @Description Store a new department in database
// @Param department body models.DepartmentRequest true "Store department"
// @Produce application/json
// @Tags department
// @Success 200 {object} models.DepartmentResponse{}
// @Router /departments [post]
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

// UpdateDepartment	godoc
// @Summary	Update department
// @Description Update a department in database
// @Param departmentID path int true "Update department by ID"
// @Param department body models.DepartmentRequest true "Update department by ID"
// @Produce application/json
// @Tags department
// @Success 200 {object} models.DepartmentResponse{}
// @Router /departments/{departmentID} [put]
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

// DestroyDepartment	godoc
// @Summary	Destroy department
// @Description Delete a department in database
// @Param departmentID path int true "Delete department by ID"
// @Produce application/json
// @Tags department
// @Success 200 {object} models.DepartmentResponse{}
// @Router /departments/{departmentID} [delete]
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
