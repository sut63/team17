package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/faculty"
)

// FacultyController defines the struct for the faculty controller
type FacultyController struct {
	client *ent.Client
	router gin.IRouter
}

// GetFaculty handles GET requests to retrieve a faculty entity
// @Summary Get a faculty entity by ID
// @Description get faculty by ID
// @ID get-faculty
// @Produce  json
// @Param id path int true "Faculty ID"
// @Success 200 {object} ent.Faculty
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facultys/{id} [get]
func (ctl *FacultyController) GetFaculty(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fa, err := ctl.client.Faculty.
		Query().
		Where(faculty.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fa)
}

// ListFaculty handles request to get a list of faculty entities
// @Summary List faculty entities
// @Description list faculty entities
// @ID list-faculty
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Faculty
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facultys [get]
func (ctl *FacultyController) ListFaculty(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 8
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 8, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 8, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	facultys, err := ctl.client.Faculty.
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

	c.JSON(200, facultys)
}

// NewFacultyController creates and registers handles for the faculty controller
func NewFacultyController(router gin.IRouter, client *ent.Client) *FacultyController {
	fac := &FacultyController{
		client: client,
		router: router,
	}

	fac.register()
	return fac

}

// FacultyController registers routes to the main engine
func (ctl *FacultyController) register() {
	facultys := ctl.router.Group("/facultys")

	// CRUD
	facultys.GET(":id", ctl.GetFaculty)
	facultys.GET("", ctl.ListFaculty)

}
