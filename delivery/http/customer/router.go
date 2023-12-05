package customer

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/customer"
)

type controller struct {
	customerService customer.Service
}

// NewCustomerController : Instance for register Customer Service
func NewCustomerController(customerService customer.Service) *controller {
	return &controller{customerService: customerService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/customer/", c.Add)
	v1.GET("/customer/:id", c.Get)
	v1.GET("/customer/", c.GetList)
	v1.PUT("/customer/:id", c.Update)
	v1.DELETE("/customer/:id", c.Delete)
}
