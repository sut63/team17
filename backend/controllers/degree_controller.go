package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ASUS/app/ent"
	"github.com/ASUS/app/ent/course"
	"github.com/ASUS/app/ent/degree"
	"github.com/gin-gonic/gin"
)

// DegreeController defines the struct for the degree controller
type DegreeController struct {
	client *ent.Client
	router gin.IRouter
}

// Degree defines the struct for the degree
type Degree struct {
	Degree string
	id     int
}

// CreateDegree handles POST requests for adding degree entities
// @Summary Create degree
// @Description Create degree
// @ID create-degree
// @Accept   json
// @Produce  json
// @Param degree body ent.Degree true "Degree entity"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [post]
func (ctl *DegreeController) CreateDegree(c *gin.Context) {
	obj := Degree{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "degree binding failed",
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(obj.id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "course not found",
		})
		return
	}
	fa, err := ctl.client.Degree.
		Create().
		SetDegree(obj.Degree).
		SetCourse(co).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, fa)
}

// GetDegree handles GET requests to retrieve a degree entity
// @Summary Get a degree entity by ID
// @Description get degree by ID
// @ID get-degree
// @Produce  json
// @Param id path int true "Degree ID"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees/{id} [get]
func (ctl *DegreeController) GetDegree(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fa, err := ctl.client.Degree.
		Query().
		WithCourse().
		Where(degree.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fa)
}

// ListDegree handles request to get a list of degree entities
// @Summary List degree entities
// @Description list degree entities
// @ID list-degree
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [get]
func (ctl *DegreeController) ListDegree(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 8
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 8, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 8, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	degrees, err := ctl.client.Degree.
		Query().
		WithCourse().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, degrees)
}

// DeleteDegree handles DELETE requests to delete a degree entity
// @Summary Delete a degree entity by ID
// @Description get degree by ID
// @ID delete-degree
// @Produce  json
// @Param id path int true "Degree ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees/{id} [delete]
func (ctl *DegreeController) DeleteDegree(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Degree.
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

// NewDegreeController creates and registers handles for the degree controller
func NewDegreeController(router gin.IRouter, client *ent.Client) *DegreeController {
	fac := &DegreeController{
		client: client,
		router: router,
	}

	fac.register()
	return fac

}

// DegreeController registers routes to the main engine
func (ctl *DegreeController) register() {
	degrees := ctl.router.Group("/degrees")

	// CRUD
	degrees.POST("", ctl.CreateDegree)
	degrees.GET(":id", ctl.GetDegree)
	degrees.GET("", ctl.ListDegree)
	degrees.DELETE("", ctl.DeleteDegree)

}
