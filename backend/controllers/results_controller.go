package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	Group     int
	Timed     string
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
	std, err := ctl.client.Student.
		Query().
		Where(student.IDEQ(int(obj.StudentID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Student not found",
		})
		return
	}
	yea, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(obj.YearID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Year not found",
		})
		return
	}
	term, err := ctl.client.Term.
		Query().
		Where(term.IDEQ(int(obj.TermID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Term not found",
		})
		return
	}
	subj, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.SubjectID))).
		Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Subject not found",
		})
		return
	}

	timereserv, err := time.Parse(time.RFC3339, obj.Timed)
	t2 := timereserv.Format("2006-01-02T15:04:05Z07:00")
	if t2 == "0001-01-01T00:00:00Z" {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "time null",
		})
		return
	}
	t, err := ctl.client.Results.
		Create().
		SetGrade(obj.Grade).
		SetResuStud(std).
		SetResuYear(yea).
		SetResuSubj(subj).
		SetResuTerm(term).
		SetTime(timereserv).
		SetGroup(obj.Group).
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
		"data":   t,
	})
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

// ListResultssomting handles request to get a list of Resultssomting entities
// @Summary List Resultssomting entities by id
// @Description list Resultssomting entities by id
// @ID list-Resultssomting
// @Produce json
// @Param year  query int false "year"
// @Param term query int false "term"
// @Param id path int true "Resultssomting ID"
// @Success 200 {array} ent.Results
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resultss111/{id} [get]
func (ctl *ResultsController) ListResultssomting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	id1, err := strconv.ParseInt(c.Query("year"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	id2, err := strconv.ParseInt(c.Query("term"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resultss, err := ctl.client.Results.
		Query().
		WithResuStud().
		WithResuYear().
		WithResuSubj().
		WithResuTerm().
		Where(results.HasResuStudWith(student.IDEQ(int(id))), results.HasResuYearWith(year.IDEQ(int(id1))), results.HasResuTermWith(term.IDEQ(int(id2)))).
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
	resultss111 := ctl.router.Group("/resultss111")

	resultss111.GET(":id", ctl.ListResultssomting)

	// CRUD
	resultss.POST("", ctl.CreateResults)
	resultss.GET(":id", ctl.GetResults)
	resultss.PUT(":id", ctl.UpdateResults)
	resultss.DELETE(":id", ctl.DeleteResults)
}
