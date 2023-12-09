package model

type Part struct {
	PartID        int     `json:"part_id"`
	PartName      string  `json:"part_name"`
	Price         float64 `json:"price"`
	SupplierID    int     `json:"supplier_id"`
	StockQuantity int     `json:"stock_quantity"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=-1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}
