package price_changes_log

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/price_changes_log"
)

type controller struct {
	pcLogService price_changes_log.Service
}

// NewPriceChangesController : Instance for register Part Service
func NewPriceChangesController(pcLog price_changes_log.Service) *controller {
	return &controller{pcLogService: pcLog}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.GET("/changes-log/:id", c.Get)
	v1.GET("/changes-log/", c.GetList)
}
