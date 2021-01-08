package controllers

import (
	"context"
	"strconv"

	"github.com/sut63/team17/app/ent/subdistrict"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
)

// SubdistrictController defines the struct for the subdistrict controller
type SubdistrictController struct {
	client *ent.Client
	router gin.IRouter
}

// GetSubdistrict handles GET requests to retrieve a subdistrict entity
// @Summary Get a subdistrict entity by ID
// @Description get subdistrict by ID
// @ID get-subdistrict
// @Produce  json
// @Param id path int true "Subdistrict ID"
// @Success 200 {object} ent.Subdistrict
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subdistricts/{id} [get]
func (ctl *SubdistrictController) GetSubdistrict(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Subdistrict.
		Query().
		Where(subdistrict.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSubdistrict handles request to get a list of subdistrict entities
// @Summary List subdistrict entities
// @Description list subdistrict entities
// @ID list-subdistrict
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subdistrict
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subdistricts [get]
func (ctl *SubdistrictController) ListSubdistrict(c *gin.Context) {
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

	histories, err := ctl.client.Subdistrict.
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

// NewSubdistrictController creates and registers handles for the subdistrict controller
func NewSubdistrictController(router gin.IRouter, client *ent.Client) *SubdistrictController {
	uc := &SubdistrictController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *SubdistrictController) register() {
	subdistricts := ctl.router.Group("/subdistricts")

	subdistricts.GET("", ctl.ListSubdistrict)

	// CRUD
	subdistricts.GET(":id", ctl.GetSubdistrict)
}
