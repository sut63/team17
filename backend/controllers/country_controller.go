package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/country"
)

// CountryController defines the struct for the country controller
type CountryController struct {
	client *ent.Client
	router gin.IRouter
}

// GetCountry handles GET requests to retrieve a country entity
// @Summary Get a country entity by ID
// @Description get country by ID
// @ID get-country
// @Produce  json
// @Param id path int true "Country ID"
// @Success 200 {object} ent.Country
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /countrys/{id} [get]
func (ctl *CountryController) GetCountry(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Country.
		Query().
		Where(country.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListCountry handles request to get a list of country entities
// @Summary List country entities
// @Description list country entities
// @ID list-country
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Country
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /countrys [get]
func (ctl *CountryController) ListCountry(c *gin.Context) {
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

	histories, err := ctl.client.Country.
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

// NewCountryController creates and registers handles for the country controller
func NewCountryController(router gin.IRouter, client *ent.Client) *CountryController {
	uc := &CountryController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *CountryController) register() {
	countrys := ctl.router.Group("/countrys")

	countrys.GET("", ctl.ListCountry)

	// CRUD
	countrys.GET(":id", ctl.GetCountry)
}
