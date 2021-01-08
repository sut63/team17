package controllers

import (
	"context"
	"strconv"

	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/agency"
	"github.com/gin-gonic/gin"
)

// AgencyController defines the struct for the agency controller
type AgencyController struct {
	client *ent.Client
	router gin.IRouter
}
// GetAgency handles GET requests to retrieve a agency entity
// @Summary Get a agency entity by ID
// @Description get agency by ID
// @ID get-agency
// @Produce  json
// @Param id path int true "Agency ID"
// @Success 200 {object} ent.Agency
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /agencys/{id} [get]
func (ctl *AgencyController) GetAgency(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	i, err := ctl.client.Agency.
		Query().
		Where(agency.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, i)
}

// ListAgency handles request to get a list of agency entities
// @Summary List agency entities
// @Description list agency entities
// @ID list-agency
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Agency
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /agencys [get]
func (ctl *AgencyController) ListAgency(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 13
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 13, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 13, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	agencys, err := ctl.client.Agency.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, agencys)
}

// NewAgencyController creates and registers handles for the agency controller
func NewAgencyController(router gin.IRouter, client *ent.Client) *AgencyController {
	ic := &AgencyController{
		client: client,
		router: router,
	}
	ic.register()
	return ic
}

// InitAgencyController registers routes to the main engine
func (ctl *AgencyController) register() {
	agencys := ctl.router.Group("/agencys")

	// CRUD
	agencys.GET(":id", ctl.GetAgency)
	agencys.GET("", ctl.ListAgency)
	
}



