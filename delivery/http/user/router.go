package user

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/user"
)

type controller struct {
	userService user.Service
}

// NewUserController : Instance for register User Service
func NewUserController(userService user.Service) *controller {
	return &controller{userService: userService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/user/", c.Add)
	v1.GET("/user/:id", c.Get)
	v1.GET("/user/", c.GetList)
	v1.PUT("/user/:id", c.Update)
	v1.DELETE("/user/:id", c.Delete)
}
