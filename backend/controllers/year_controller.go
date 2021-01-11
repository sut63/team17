package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/year"
)

// YearController defines the struct for the year controller
type YearController struct {
	client *ent.Client
	router gin.IRouter
}

// GetYear handles GET requests to retrieve a year entity
// @Summary Get a year entity by ID
// @Description get year by ID
// @ID get-year
// @Produce  json
// @Param id path int true "Year ID"
// @Success 200 {object} ent.Year
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years/{id} [get]
func (ctl *YearController) GetYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListYear handles request to get a list of year entities
// @Summary List year entities
// @Description list year entities
// @ID list-year
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Year
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years [get]
func (ctl *YearController) ListYear(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 100
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

	years, err := ctl.client.Year.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, years)
}

// NewYearController creates and registers handles for the year controller
func NewYearController(router gin.IRouter, client *ent.Client) *YearController {
	ggy := &YearController{
		client: client,
		router: router,
	}
	ggy.register()
	return ggy
}

// InitYearController registers routes to the main engine
func (ctl *YearController) register() {
	years := ctl.router.Group("/years")
	years.GET("", ctl.ListYear)
	// CRUD
	years.GET(":id", ctl.GetYear)
}
