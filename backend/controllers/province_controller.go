package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/continent"
	"github.com/sut63/team17/app/ent/country"
	"github.com/sut63/team17/app/ent/region"
)

type ProvinceController struct {
	client *ent.Client
	router gin.IRouter
}

type Province struct {
	Province    string
	District    string
	Subdistrict string
	Postal      int
	Region      int
	Continent   int
	Country     int
}

// CreateProvince handles POST requests for adding province entities
// @Summary Create province
// @Description Create province
// @ID create-province
// @Accept   json
// @Produce  json
// @Param province body Province true "Province entity"
// @Success 200 {object} ent.Province
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /provinces [post]
func (ctl *ProvinceController) ProvinceCreate(c *gin.Context) {
	obj := Province{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "province binding failed",
		})
		return
	}

	cu, err := ctl.client.Country.
		Query().
		Where(country.IDEQ(int(obj.Country))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "country not found",
		})
		return
	}

	cn, err := ctl.client.Continent.
		Query().
		Where(continent.IDEQ(int(obj.Continent))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "continent not found",
		})
		return
	}

	re, err := ctl.client.Region.
		Query().
		Where(region.IDEQ(int(obj.Region))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "region not found",
		})
		return
	}

	save, err := ctl.client.Province.
		Create().
		SetProvince(obj.Province).
		SetDistrict(obj.District).
		SetSubdistrict(obj.Subdistrict).
		SetPostal(obj.Postal).
		SetProvCoun(cu).
		SetProvCont(cn).
		SetProvRegi(re).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   save,
	})
}

// ListProvince handles request to get a list of province entities
// @Summary List province entities
// @Description list province entities
// @ID list-province
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Province
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /provinces [get]
func (ctl *ProvinceController) ListProvince(c *gin.Context) {
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

	provinces, err := ctl.client.Province.
		Query().
		WithProvCoun().
		WithProvCont().
		WithProvRegi().
		WithProvStud().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, provinces)
}

// NewProvinceController creates and registers handles for the province controller
func NewProvinceController(router gin.IRouter, client *ent.Client) *ProvinceController {
	pvc := &ProvinceController{
		client: client,
		router: router,
	}
	pvc.register()
	return pvc
}

// InitProvinceController registers routes to the main engine
func (ctl *ProvinceController) register() {
	provinces := ctl.router.Group("/provinces")

	// CRUD
	provinces.POST("", ctl.ProvinceCreate)
	provinces.GET("", ctl.ListProvince)
}
