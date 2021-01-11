package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/subject"
)

// SubjectController defines the struct for the subject controller
type SubjectController struct {
	client *ent.Client
	router gin.IRouter
}

// GetSubject handles GET requests to retrieve a subject entity
// @Summary Get a subject entity by ID
// @Description get subject by ID
// @ID get-subject
// @Produce  json
// @Param id path int true "Subject ID"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects/{id} [get]
func (ctl *SubjectController) GetSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSubject handles request to get a list of subject entities
// @Summary List subject entities
// @Description list subject entities
// @ID list-subject
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects [get]
func (ctl *SubjectController) ListSubject(c *gin.Context) {
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

	subjects, err := ctl.client.Subject.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, subjects)
}

// NewSubjectController creates and registers handles for the subject controller
func NewSubjectController(router gin.IRouter, client *ent.Client) *SubjectController {
	ggs := &SubjectController{
		client: client,
		router: router,
	}
	ggs.register()
	return ggs
}

// InitSubjectController registers routes to the main engine
func (ctl *SubjectController) register() {
	subjects := ctl.router.Group("/subjects")

	subjects.GET("", ctl.ListSubject)

	// CRUD

	subjects.GET(":id", ctl.GetSubject)

}
