package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/grazi/app/ent"
	"github.com/grazi/app/ent/term"
)

// TermController defines the struct for the term controller
type TermController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTerm handles POST requests for adding term entities
// @Summary Create term
// @Description Create term
// @ID create-term
// @Accept   json
// @Produce  json
// @Param term body ent.Term true "Term entity"
// @Success 200 {object} ent.Term
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terms [post]
func (ctl *TermController) CreateTerm(c *gin.Context) {
	obj := ent.Term{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "term binding failed",
		})
		return
	}

	u, err := ctl.client.Term.
		Create().
		SetSemester(obj.semester).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
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

// DeleteTerm handles DELETE requests to delete a term entity
// @Summary Delete a term entity by ID
// @Description get term by ID
// @ID delete-term
// @Produce  json
// @Param id path int true "Term ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terms/{id} [delete]
func (ctl *TermController) DeleteTerm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Term.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTerm handles PUT requests to update a term entity
// @Summary Update a term entity by ID
// @Description update term by ID
// @ID update-term
// @Accept   json
// @Produce  json
// @Param id path int true "Term ID"
// @Param term body ent.Term true "Term entity"
// @Success 200 {object} ent.Term
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terms/{id} [put]
func (ctl *TermController) UpdateTerm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Term{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "term binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Term.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
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
	terms.POST("", ctl.CreateTerm)
	terms.GET(":id", ctl.GetTerm)
	terms.PUT(":id", ctl.UpdateTerm)
	terms.DELETE(":id", ctl.DeleteTerm)
}
