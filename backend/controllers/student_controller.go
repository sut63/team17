package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/gender"
	"github.com/sut63/team17/app/ent/student"
)

type StudentController struct {
	client *ent.Client
	router gin.IRouter
}

//Student struct
type Student struct {
	Fname  string
	Lname  string
	School string
	Addr   string
	Email  string
	Tel    int
	Sex    int
}

// CreateStudent handles POST requests for adding student entities
// @Summary Create student
// @Description Create student
// @ID create-student
// @Accept   json
// @Produce  json
// @Param student body Student true "Student entity"
// @Success 200 {object} ent.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students [post]
func (ctl *StudentController) CreateStudent(c *gin.Context) {
	obj := Student{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "student binding failed",
		})
		return
	}
	s, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.sex))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	save, err := ctl.client.Student.
		Create().
		SetFname(obj.fname).
		SetLname(obj.lname).
		SetRecentAddress(obj.addr).
		SetSchool(obj.school).
		SetEmail(obj.email).
		SetTelephone(obj.tel).
		SetStudGend(s).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, save)
}

// GetStudent handles GET requests to retrieve a student entity
// @Summary Get a student entity by ID
// @Description get student by ID
// @ID get-student
// @Produce  json
// @Param id path int true "Student ID"
// @Success 200 {object} ent.Student
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students/{id} [get]
func (ctl *StudentController) GetStudent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Student.
		Query().
		Where(student.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListStudent handles request to get a list of student entities
// @Summary List student entities
// @Description list student entities
// @ID list-student
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students [get]
func (ctl *StudentController) ListStudent(c *gin.Context) {
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

	studentss, err := ctl.client.Student.
		Query().
		WithStudGend().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, studentss)
}

// NewStudentController creates and registers handles for the student controller
func NewStudentController(router gin.IRouter, client *ent.Client) *StudentController {
	pvc := &StudentController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *StudentController) register() {
	student := ctl.router.Group("/students")

	student.POST("", ctl.CreateStudent)
	student.GET("", ctl.ListStudent)
	student.GET(":id", ctl.GetStudent)

}
