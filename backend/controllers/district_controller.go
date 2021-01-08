package controllers

import (
	"context"
	"strconv"

	"github.com/sut63/team17/app/ent/district"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
)

// DistrictController defines the struct for the district controller
type DistrictController struct {
	client *ent.Client
	router gin.IRouter
}

// GetDistrict handles GET requests to retrieve a district entity
// @Summary Get a district entity by ID
// @Description get district by ID
// @ID get-district
// @Produce  json
// @Param id path int true "District ID"
// @Success 200 {object} ent.District
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /districts/{id} [get]
func (ctl *DistrictController) GetDistrict(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.District.
		Query().
		Where(district.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListDistrict handles request to get a list of district entities
// @Summary List district entities
// @Description list district entities
// @ID list-district
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.District
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /districts [get]
func (ctl *DistrictController) ListDistrict(c *gin.Context) {
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

	histories, err := ctl.client.District.
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

// NewDistrictController creates and registers handles for the district controller
func NewDistrictController(router gin.IRouter, client *ent.Client) *DistrictController {
	uc := &DistrictController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *DistrictController) register() {
	districts := ctl.router.Group("/districts")

	districts.GET("", ctl.ListDistrict)

	// CRUD
	districts.GET(":id", ctl.GetDistrict)
}
