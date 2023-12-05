package part

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/part"
)

type controller struct {
	partService part.Service
}

// NewPartController : Instance for register Part Service
func NewPartController(partService part.Service) *controller {
	return &controller{partService: partService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/part/", c.Add)
	v1.GET("/part/:id", c.Get)
	v1.GET("/part/", c.GetList)
	v1.PUT("/part/:id", c.Update)
	v1.DELETE("/part/:id", c.Delete)
}
