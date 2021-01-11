package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/term"
)

// TermController defines the struct for the term controller
type TermController struct {
	client *ent.Client
	router gin.IRouter
}

// GetTerm handles GET requests to retrieve a term entity
// @Summary Get a term entity by ID
// @Description get term by ID
// @ID get-term
// @Produce  json
// @Param id path int true "Term ID"
// @Success 200 {object} ent.Term
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terms/{id} [get]
func (ctl *TermController) GetTerm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Term.
		Query().
		Where(term.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListTerm handles request to get a list of term entities
// @Summary List term entities
// @Description list term entities
// @ID list-term
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Term
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terms [get]
func (ctl *TermController) ListTerm(c *gin.Context) {
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

	terms, err := ctl.client.Term.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, terms)
}

// NewTermController creates and registers handles for the term controller
func NewTermController(router gin.IRouter, client *ent.Client) *TermController {
	ggt := &TermController{
		client: client,
		router: router,
	}
	ggt.register()
	return ggt
}

// InitTermController registers routes to the main engine
func (ctl *TermController) register() {
	terms := ctl.router.Group("/terms")

	terms.GET("", ctl.ListTerm)

	// CRUD

	terms.GET(":id", ctl.GetTerm)

}
