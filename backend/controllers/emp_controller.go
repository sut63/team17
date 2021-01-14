package controllers

import (
	"context"
	"strconv"

	"github.com/sut63/team17/app/ent/emp"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
)

// EmpController defines the struct for the emp controller
type EmpController struct {
	client *ent.Client
	router gin.IRouter
}

// GetEmp handles GET requests to retrieve a emp entity
// @Summary Get a emp entity by ID
// @Description get emp by ID
// @ID get-emp
// @Produce  json
// @Param id path int true "Emp ID"
// @Success 200 {object} ent.Emp
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /emps/{id} [get]
func (ctl *EmpController) GetEmp(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Emp.
		Query().
		Where(emp.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListEmp handles request to get a list of emp entities
// @Summary List emp entities
// @Description list emp entities
// @ID list-emp
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Emp
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /emps [get]
func (ctl *EmpController) ListEmp(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	histories, err := ctl.client.Emp.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, histories)
}

// NewEmpController creates and registers handles for the emp controller
func NewEmpController(router gin.IRouter, client *ent.Client) *EmpController {
	uc := &EmpController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *EmpController) register() {
	emps := ctl.router.Group("/emps")

	emps.GET("", ctl.ListEmp)

	// CRUD
	emps.GET(":id", ctl.GetEmp)
}
