package controllers

import (
	"context"
	"strconv"

	"github.com/jeeninee/app/ent/postal"

	"github.com/gin-gonic/gin"
	"github.com/jeeninee/app/ent"
)

// PostalController defines the struct for the postal controller
type PostalController struct {
	client *ent.Client
	router gin.IRouter
}

// GetPostal handles GET requests to retrieve a postal entity
// @Summary Get a postal entity by ID
// @Description get postal by ID
// @ID get-postal
// @Produce  json
// @Param id path int true "Postal ID"
// @Success 200 {object} ent.Postal
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /postals/{id} [get]
func (ctl *PostalController) GetPostal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Postal.
		Query().
		Where(postal.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListPostal handles request to get a list of postal entities
// @Summary List postal entities
// @Description list postal entities
// @ID list-postal
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Postal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /postals [get]
func (ctl *PostalController) ListPostal(c *gin.Context) {
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

	histories, err := ctl.client.Postal.
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

// NewPostalController creates and registers handles for the postal controller
func NewPostalController(router gin.IRouter, client *ent.Client) *PostalController {
	uc := &PostalController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *PostalController) register() {
	postals := ctl.router.Group("/postals")

	postals.GET("", ctl.ListPostal)

	// CRUD
	postals.GET(":id", ctl.GetPostal)
}
