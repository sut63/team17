package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/se63/team17/app/ent"
	"github.com/se63/team17/app/ent/course"
	"github.com/se63/team17/app/ent/institution"
)

// InstitutionController defines the struct for the institution controller
type InstitutionController struct {
	client *ent.Client
	router gin.IRouter
}

//Institution defines the struct for the institution
type Institution struct {
	Institution string
	id          int
}

// CreateInstitution handles POST requests for adding institution entities
// @Summary Create institution
// @Description Create institution
// @ID create-institution
// @Accept   json
// @Produce  json
// @Param institution body ent.Institution true "Institution entity"
// @Success 200 {object} ent.Institution
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /institutions [post]
func (ctl *InstitutionController) CreateInstitution(c *gin.Context) {
	obj := Institution{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "institution binding failed",
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(obj.id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "course not found",
		})
		return
	}
	fa, err := ctl.client.Institution.
		Create().
		SetInstitution(obj.Institution).
		SetCourse(co).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, fa)
}

// GetInstitution handles GET requests to retrieve a institution entity
// @Summary Get a institution entity by ID
// @Description get institution by ID
// @ID get-institution
// @Produce  json
// @Param id path int true "Institution ID"
// @Success 200 {object} ent.Institution
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /institutions/{id} [get]
func (ctl *InstitutionController) GetInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fa, err := ctl.client.Institution.
		Query().
		WithCourse().
		Where(institution.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fa)
}

// ListInstitution handles request to get a list of institution entities
// @Summary List institution entities
// @Description list institutionentities
// @ID list-institution
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Institution
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /institutions [get]
func (ctl *InstitutionController) ListInstitution(c *gin.Context) {
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

	institutions, err := ctl.client.Institution.
		Query().
		WithCourse().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, institutions)
}

// DeleteInstitution handles DELETE requests to delete a institution entity
// @Summary Delete a institution entity by ID
// @Description get institution by ID
// @ID delete-institution
// @Produce  json
// @Param id path int true "Institution ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /institutions/{id} [delete]
func (ctl *InstitutionController) DeleteInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Institution.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewInstitutionController creates and registers handles for the institution controller
func NewInstitutionController(router gin.IRouter, client *ent.Client) *InstitutionController {
	fac := &InstitutionController{
		client: client,
		router: router,
	}

	fac.register()
	return fac

}

// InstitutionController registers routes to the main engine
func (ctl *InstitutionController) register() {
	institutions := ctl.router.Group("/institutions")

	// CRUD
	institutions.POST("", ctl.CreateInstitution)
	institutions.GET(":id", ctl.GetInstitution)
	institutions.GET("", ctl.ListInstitution)
	institutions.DELETE("", ctl.DeleteInstitution)

}
