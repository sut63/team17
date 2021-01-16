package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/institution"
)

// InstitutionController defines the struct for the institution controller
type InstitutionController struct {
	client *ent.Client
	router gin.IRouter
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
	limit := 100
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 23, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 23, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	institutions, err := ctl.client.Institution.
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

	c.JSON(200, institutions)
}

// NewInstitutionController creates and registers handles for the institution controller
func NewInstitutionController(router gin.IRouter, client *ent.Client) *InstitutionController {
	ggt := &InstitutionController{
		client: client,
		router: router,
	}
	ggt.register()
	return ggt
}

// InstitutionController registers routes to the main engine
func (ctl *InstitutionController) register() {
	institutions := ctl.router.Group("/institutions")

	// CRUD
	institutions.GET(":id", ctl.GetInstitution)
	institutions.GET("", ctl.ListInstitution)

}
