package controllers

import (
	"context"
	"strconv"

	"github.com/jeeninee/app/ent/region"

	"github.com/gin-gonic/gin"
	"github.com/jeeninee/app/ent"
)

// RegionController defines the struct for the region controller
type RegionController struct {
	client *ent.Client
	router gin.IRouter
}

// GetRegion handles GET requests to retrieve a region entity
// @Summary Get a region entity by ID
// @Description get region by ID
// @ID get-region
// @Produce  json
// @Param id path int true "Region ID"
// @Success 200 {object} ent.Region
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /regions/{id} [get]
func (ctl *RegionController) GetRegion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Region.
		Query().
		Where(region.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRegion handles request to get a list of region entities
// @Summary List region entities
// @Description list region entities
// @ID list-region
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Region
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /regions [get]
func (ctl *RegionController) ListRegion(c *gin.Context) {
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

	histories, err := ctl.client.Region.
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

// NewRegionController creates and registers handles for the region controller
func NewRegionController(router gin.IRouter, client *ent.Client) *RegionController {
	uc := &RegionController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *RegionController) register() {
	regions := ctl.router.Group("/regions")

	regions.GET("", ctl.ListRegion)

	// CRUD
	regions.GET(":id", ctl.GetRegion)
}
