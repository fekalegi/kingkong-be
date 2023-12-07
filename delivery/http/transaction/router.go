package transaction

import (
	"github.com/gin-gonic/gin"
	"kingkong-be/domain/transaction"
)

type controller struct {
	transactionService transaction.Service
}

// NewTransactionController : Instance for register Transaction Service
func NewTransactionController(transactionService transaction.Service) *controller {
	return &controller{transactionService: transactionService}
}

func (c *controller) Route(e *gin.RouterGroup) {
	v1 := e.Group("/v1")
	v1.POST("/transaction/", c.Add)
	v1.GET("/transaction/:id", c.Get)
	v1.GET("/transaction/", c.GetList)
	v1.PUT("/transaction/:id", c.Update)
	v1.DELETE("/transaction/:id", c.Delete)
}
