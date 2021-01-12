package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/continent"
)

// ContinentController defines the struct for the continent controller
type ContinentController struct {
	client *ent.Client
	router gin.IRouter
}

// GetContinent handles GET requests to retrieve a continent entity
// @Summary Get a continent entity by ID
// @Description get continent by ID
// @ID get-continent
// @Produce  json
// @Param id path int true "Continent ID"
// @Success 200 {object} ent.Continent
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /continents/{id} [get]
func (ctl *ContinentController) GetContinent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Continent.
		Query().
		Where(continent.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListContinent handles request to get a list of continent entities
// @Summary List continent entities
// @Description list continent entities
// @ID list-continent
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Continent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /continents [get]
func (ctl *ContinentController) ListContinent(c *gin.Context) {
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

	histories, err := ctl.client.Continent.
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

// NewContinentController creates and registers handles for the continent controller
func NewContinentController(router gin.IRouter, client *ent.Client) *ContinentController {
	uc := &ContinentController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *ContinentController) register() {
	continents := ctl.router.Group("/continents")

	continents.GET("", ctl.ListContinent)

	// CRUD
	continents.GET(":id", ctl.GetContinent)
}
