package supplier

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/supplier"
)

type controller struct {
	supplierService supplier.Service
}

// NewSupplierController : Instance for register Supplier Service
func NewSupplierController(supplierService supplier.Service) *controller {
	return &controller{supplierService: supplierService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/supplier/", c.Add)
	v1.GET("/supplier/:id", c.Get)
	v1.GET("/supplier/", c.GetList)
	v1.PUT("/supplier/:id", c.Update)
	v1.DELETE("/supplier/:id", c.Delete)
}
