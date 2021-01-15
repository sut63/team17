package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/agency"
	"github.com/sut63/team17/app/ent/place"
	"github.com/sut63/team17/app/ent/student"
	"github.com/sut63/team17/app/ent/term"
	"github.com/sut63/team17/app/ent/year"
)

// ActivityController defines the struct for the activity controller
type ActivityController struct {
	client *ent.Client
	router gin.IRouter
}

// Activity defines the struct for the activity
type Activity struct {
	Agency       int
	Place        int
	Added        string
	Hours        string
	Activityname string
	Year         int
	Student      int
	Term         int
}

// CreateActivity handles POST requests for adding activity entities
// @Summary Create activity
// @Description Create activity
// @ID create-activity
// @Accept   json
// @Produce  json
// @Param activity body Activity true "Activity entity"
// @Success 200 {object} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /activitys [post]
func (ctl *ActivityController) CreateActivity(c *gin.Context) {
	obj := Activity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "activity binding failed",
		})
		return

	}

	ag, err := ctl.client.Agency.
		Query().
		Where(agency.IDEQ(int(obj.Agency))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "agency not found",
		})
		return
	}
	p, err := ctl.client.Place.
		Query().
		Where(place.IDEQ(int(obj.Place))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "place not found",
		})
		return
	}
	y, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(obj.Year))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "year not found",
		})
		return
	}
	st, err := ctl.client.Student.
		Query().
		Where(student.IDEQ(int(obj.Student))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "student not found",
		})
		return
	}
	t, err := ctl.client.Term.
		Query().
		Where(term.IDEQ(int(obj.Term))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "term not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)

	save, err := ctl.client.Activity.
		Create().
		SetACTIVITYNAME(obj.Activityname).
		SetAdded(time).
		SetActiAgen(ag).
		SetActiPlace(p).
		SetHours(obj.Hours).
		SetActiYear(y).
		SetActiStud(st).
		SetActiTerm(t).
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

// DeleteActivity handles DELETE requests to delete a activity entity
// @Summary Delete a activity entity by ID
// @Description get activity by ID
// @ID delete-activity
// @Produce  json
// @Param id path int true "Activity ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /activitys/{id} [delete]
func (ctl *ActivityController) DeleteActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Activity.
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

// ListActivity handles request to get a list of activity entities
// @Summary List activity entities
// @Description list activity entities
// @ID list-activity
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Activity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /activitys [get]
func (ctl *ActivityController) ListActivity(c *gin.Context) {
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

	activitys, err := ctl.client.Activity.
		Query().
		WithActiAgen().
		WithActiPlace().
		WithActiYear().
		WithActiStud().
		WithActiTerm().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, activitys)
}

// NewActivityController creates and registers handles for the activity controller
func NewActivityController(router gin.IRouter, client *ent.Client) *ActivityController {
	ac := &ActivityController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

// InitActivityController registers routes to the main engine
func (ctl *ActivityController) register() {
	activitys := ctl.router.Group("/activitys")
	activitys.GET("", ctl.ListActivity)

	// CRUD
	activitys.POST("", ctl.CreateActivity)
	activitys.DELETE(":id", ctl.DeleteActivity)

}
