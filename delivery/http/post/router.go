package post

import (
	"github.com/gin-gonic/gin"
	"sharing-vision-2021/domain/post"
)

type controller struct {
	postService post.Service
}

// NewPostController : Instance for register Post Service
func NewPostController(postService post.Service) *controller {
	return &controller{postService: postService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/article/", c.Add)
	v1.GET("/article/:id", c.Get)
	v1.GET("/article/", c.GetList)
	v1.PUT("/article/:id", c.Update)
	v1.DELETE("/article/:id", c.Delete)
}
