package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/se63/team17/app/ent/course"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/faculty"
	"github.com/sut63/team17/app/ent/institution"
)

// CourseController defines the struct for the course controller
type CourseController struct {
	client *ent.Client
	router gin.IRouter
}

// Course defines the struct for the course
type Course struct {
	Coursename  string
	id          int
	Degree      int
	Institution int
	Faculty     int
}

// CreateCourse handles POST requests for adding course entities
// @Summary Create course
// @Description Create course
// @ID create-course
// @Accept   json
// @Produce  json
// @Param course body ent.Course true "Course entity"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [post]
func (ctl *CourseController) CreateCourse(c *gin.Context) {
	obj := Course{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "course binding failed",
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

	de, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(obj.Degree))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "degree not found",
		})
		return
	}

	in, err := ctl.client.Institution.
		Query().
		Where(institution.IDEQ(int(obj.Institution))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "institution not found",
		})
		return
	}

	save, err := ctl.client.Course.
		Create().
		Setcourse(obj.Coursename).
		Setfaculty(fa).
		Setdegree(de).
		Setinstitution(in).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, save)
}

// GetCourse handles GET requests to retrieve a course entity
// @Summary Get a course entity by ID
// @Description get course by ID
// @ID get-course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [get]
func (ctl *CourseController) GetCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, co)
}

// ListCourse handles request to get a list of course entities
// @Summary List course entities
// @Description list course entities
// @ID list-course
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [get]
func (ctl *CourseController) ListCourse(c *gin.Context) {
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

	courses, err := ctl.client.Course.
		Query().
		WithFaculty().
		WithDegree().
		WithInstitution().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, courses)
}

// DeleteCourse handles DELETE requests to delete a course entity
// @Summary Delete a course entity by ID
// @Description get course by ID
// @ID delete-course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [delete]
func (ctl *CourseController) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Course.
		DeleteOneID(int(id)).
		Coec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewCourseController creates and registers handles for the course controller
func NewCourseController(router gin.IRouter, client *ent.Client) *CourseController {
	coc := &CourseController{
		client: client,
		router: router,
	}

	coc.register()
	return coc

}

// InitCourseController registers routes to the main engine
func (ctl *CourseController) register() {
	courses := ctl.router.Group("/courses")

	//CRUD
	courses.POST("", ctl.CreateCourse)
	courses.GET(":id", ctl.GetCourse)
	courses.GET("", ctl.ListCourse)
	courses.DELETE("", ctl.DeleteCourse)

}
