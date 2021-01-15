package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/faculty"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/professorship"
)

type ProfessorController struct {
	client *ent.Client
	router gin.IRouter
}

type Professor struct {
	Tel           string
	Email         string
	Name          string
	Prefix        int
	Faculty       int
	Professorship int
}

// CreateProfessor handles POST requests for adding professor entities
// @Summary Create professor
// @Description Create professor
// @ID create-professor
// @Accept   json
// @Produce  json
// @Param professor body Professor true "Professor entity"
// @Success 200 {object} ent.Professor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /professors [post]
func (ctl *ProfessorController) CreateProfessor(c *gin.Context) {
	obj := Professor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "professor binding failed",
		})
		return
	}

	pf, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(obj.Prefix))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "prefix not found",
		})
		return
	}

	fa, err := ctl.client.Faculty.
		Query().
		Where(faculty.IDEQ(int(obj.Faculty))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "faculty not found",
		})
		return
	}

	ps, err := ctl.client.Professorship.
		Query().
		Where(professorship.IDEQ(int(obj.Professorship))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "professorship not found",
		})
		return
	}

	save, err := ctl.client.Professor.
		Create().
		SetName(obj.Name).
		SetEmail(obj.Email).
		SetTel(obj.Tel).
		SetProfPre(pf).
		SetProfFacu(fa).
		SetProfPros(ps).
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

// ListProfessor handles request to get a list of professor entities
// @Summary List professor entities
// @Description list professor entities
// @ID list-professor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Professor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /professors [get]
func (ctl *ProfessorController) ListProfessor(c *gin.Context) {
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

	professors, err := ctl.client.Professor.
		Query().
		WithProfPre().
		WithProfFacu().
		WithProfPros().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, professors)
}

// DeleteProfessor handles DELETE requests to delete a professor entity
// @Summary Delete a professor entity by ID
// @Description get professor by ID
// @ID delete-professor
// @Produce  json
// @Param id path int true "Professor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /professors/{id} [delete]
func (ctl *ProfessorController) DeleteProfessor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Professor.
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

// NewProfessorController creates and registers handles for the professor controller
func NewProfessorController(router gin.IRouter, client *ent.Client) *ProfessorController {
	pc := &ProfessorController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *ProfessorController) register() {
	professors := ctl.router.Group("/professors")

	professors.POST("", ctl.CreateProfessor)
	professors.GET("", ctl.ListProfessor)
	professors.DELETE(":id", ctl.DeleteProfessor)

}
