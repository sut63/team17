package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/results"
	"github.com/sut63/team17/app/ent/year"

	"github.com/sut63/team17/app/ent/student"
	"github.com/sut63/team17/app/ent/subject"
	"github.com/sut63/team17/app/ent/term"
)

// ResultsController defines the struct for the results controller
type ResultsController struct {
	client *ent.Client
	router gin.IRouter
}

// Results struct eiei
type Results struct {
	Grade     float64
	StudentID int
	YearID    int
	SubjectID int
	TermID    int
}

// CreateResults handles POST requests for adding results entities
// @Summary Create results
// @Description Create results
// @ID create-results
// @Accept   json
// @Produce  json
// @Param results body Results true "Results entity"
// @Success 200 {object} ent.Results
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss [post]
func (ctl *ResultsController) CreateResults(c *gin.Context) {
	obj := Results{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "results binding failed",
		})
		return
	}

	yea, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(obj.YearID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	subj, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.SubjectID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	std, err := ctl.client.Student.
		Query().
		Where(student.IDEQ(int(obj.StudentID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	term, err := ctl.client.Term.
		Query().
		Where(term.IDEQ(int(obj.TermID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t, err := ctl.client.Results.
		Create().
		SetGrade(obj.Grade).
		SetResuStud(std).
		SetResuYear(yea).
		SetResuSubj(subj).
		SetResuTerm(term).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetResults handles GET requests to retrieve a results entity
// @Summary Get a results entity by ID
// @Description get results by ID
// @ID get-results
// @Produce  json
// @Param id path int true "Results ID"
// @Success 200 {object} ent.Results
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss/{id} [get]
func (ctl *ResultsController) GetResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Results.
		Query().
		Where(results.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListResults handles request to get a list of results entities
// @Summary List results entities
// @Description list results entities
// @ID list-results
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Results
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss [get]
func (ctl *ResultsController) ListResults(c *gin.Context) {
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

	resultss, err := ctl.client.Results.
		Query().
		WithResuStud().
		WithResuYear().
		WithResuSubj().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resultss)
}

// DeleteResults handles DELETE requests to delete a results entity
// @Summary Delete a results entity by ID
// @Description get results by ID
// @ID delete-results
// @Produce  json
// @Param id path int true "Results ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss/{id} [delete]
func (ctl *ResultsController) DeleteResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Results.
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

// UpdateResults handles PUT requests to update a results entity
// @Summary Update a results entity by ID
// @Description update results by ID
// @ID update-results
// @Accept   json
// @Produce  json
// @Param id path int true "Results ID"
// @Param results body ent.Results true "Results entity"
// @Success 200 {object} ent.Results
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss/{id} [put]
func (ctl *ResultsController) UpdateResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Results{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "results binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Results.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewResultsController creates and registers handles for the results controller
func NewResultsController(router gin.IRouter, client *ent.Client) *ResultsController {
	ggr := &ResultsController{
		client: client,
		router: router,
	}
	ggr.register()
	return ggr
}

// InitResultsController registers routes to the main engine
func (ctl *ResultsController) register() {
	resultss := ctl.router.Group("/resultss")

	resultss.GET("", ctl.ListResults)

	// CRUD
	resultss.POST("", ctl.CreateResults)
	resultss.GET(":id", ctl.GetResults)
	resultss.PUT(":id", ctl.UpdateResults)
	resultss.DELETE(":id", ctl.DeleteResults)
}
