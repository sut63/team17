package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/continent"
	"github.com/sut63/team17/app/ent/country"
	"github.com/sut63/team17/app/ent/province"
	"github.com/sut63/team17/app/ent/region"
)

type ProvinceController struct {
	client *ent.Client
	router gin.IRouter
}

//Province struct
type Province struct {
	Prov string
	Dist string
	Subd string
	Post string
	Regi int
	Coun int
	Cont int
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
func (ctl *ProvinceController) CreateProvince(c *gin.Context) {
	obj := Province{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "province binding failed",
		})
		return
	}
	count, err := ctl.client.Country.
		Query().
		Where(country.IDEQ(int(obj.Coun))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}
	conti, err := ctl.client.Continent.
		Query().
		Where(continent.IDEQ(int(obj.Cont))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}
	regio, err := ctl.client.Region.
		Query().
		Where(region.IDEQ(int(obj.Regi))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	save, err := ctl.client.Province.
		Create().
		SetDistrict(obj.Dist).
		SetSubdistrict(obj.Subd).
		SetPostal(obj.Post).
		SetProvince(obj.Prov).
		SetProvCont(conti).
		SetProvCoun(count).
		SetProvRegi(regio).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   save,
	})
}

// GetProvince handles GET requests to retrieve a province entity
// @Summary Get a province entity by ID
// @Description get province by ID
// @ID get-province
// @Produce  json
// @Param id path int true "Province ID"
// @Success 200 {object} ent.Province
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /provinces/{id} [get]
func (ctl *ProvinceController) GetProvince(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Province.
		Query().
		Where(province.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
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

	provincess, err := ctl.client.Province.
		Query().
		WithProvCont().
		WithProvCoun().
		WithProvRegi().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, provincess)
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

func (ctl *ProvinceController) register() {
	province := ctl.router.Group("/provinces")
	province.GET("", ctl.ListProvince)
	// CRUD
	province.POST("", ctl.CreateProvince)
	province.GET(":id", ctl.GetProvince)

}
