package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/prefix"
)

type PrefixController struct {
	client *ent.Client
	router gin.IRouter
}

// GetPrefix handles GET requests to retrieve a prefix entity
// @Summary Get a prefix entity by ID
// @Description get prefix by ID
// @ID get-prefix
// @Produce  json
// @Param id path int true "Prefix ID"
// @Success 200 {object} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs/{id} [get]
func (ctl *PrefixController) GetPrefix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pf, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pf)
}

// ListPrefix handles request to get a list of prefix entities
// @Summary List prefix entities
// @Description list prefix entities
// @ID list-prefix
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs [get]
func (ctl *PrefixController) ListPrefix(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 3
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 3, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 3, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	prefixs, err := ctl.client.Prefix.
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

	c.JSON(200, prefixs)
}

// NewPrefixController creates and registers handles for the prefix controller
func NewPrefixController(router gin.IRouter, client *ent.Client) *PrefixController {
	pfc := &PrefixController{
		client: client,
		router: router,
	}

	pfc.register()

	return pfc

}

func (ctl *PrefixController) register() {
	prefixs := ctl.router.Group("/prefixs")

	prefixs.GET(":id", ctl.GetPrefix)
	prefixs.GET("", ctl.ListPrefix)

}
