package handlers

import (
	"github.com/gin-gonic/gin"

	"goContextDiscovery/cmd/internal/core/domain"
	"goContextDiscovery/cmd/internal/core/ports"
)

type HttpHandler struct {
	personService ports.PersonService
}

func NewHttpHandler(service ports.PersonService) *HttpHandler {
	return &HttpHandler{
		personService: service,
	}
}

func (h *HttpHandler) Find(c *gin.Context) {
	person, err := h.personService.Find(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, person)
}

func (h *HttpHandler) SignUp(c *gin.Context) {
	var body domain.Person
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return
	}

	err := h.personService.SignUp(body.Name, body.Age)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(202, nil)
}
