package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team17/app/ent"
	"github.com/sut63/team17/app/ent/professorship"
)
type ProfessorshipController struct {
	client *ent.Client
	router gin.IRouter
}


// GetProfessorship handles GET requests to retrieve a professorship entity
// @Summary Get a professorship entity by ID
// @Description get professorship by ID
// @ID get-professorship
// @Produce  json
// @Param id path int true "Professorship ID"
// @Success 200 {object} ent.Professorship
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /professorships/{id} [get]
func (ctl *ProfessorshipController) GetProfessorship(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ps, err := ctl.client.Professorship.
		Query().
		Where(professorship.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ps)
}

// ListProfessorship handles request to get a list of professorship entities
// @Summary List professorship entities
// @Description list professorship entities
// @ID list-professorship
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Professorship
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /professorships [get]
func (ctl *ProfessorshipController) ListProfessorship(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 4
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 4, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 4, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	professorships, err := ctl.client.Professorship.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, professorships)
}

// NewProfessorshipController creates and registers handles for the professorship controller
func NewProfessorshipController(router gin.IRouter, client *ent.Client) *ProfessorshipController {
	psc := &ProfessorshipController{
		client: client,
		router: router,
	}

	psc.register()

	return psc

}

func (ctl *ProfessorshipController) register() {
	professorships := ctl.router.Group("/professorships")

	professorships.GET(":id", ctl.GetProfessorship)
	professorships.GET("", ctl.ListProfessorship)

}
