package part

import (
	"kingkong-be/delivery/http/part/model"
	entity "kingkong-be/domain/part"
)

func mapRequestAddPart(req *model.Part, e *entity.Part) {
	e.PartID = req.PartID
	e.PartName = req.PartName
	e.Price = req.Price
	e.StockQuantity = req.StockQuantity
}
