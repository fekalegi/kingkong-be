package part

import (
	"time"
)

type Part struct {
	PartID        int        `json:"part_id" gorm:"primaryKey;autoIncrement"`
	SupplierID    int        `json:"supplier_id"`
	SupplierName  string     `json:"supplier_name" gorm:"->"`
	PartName      string     `json:"part_name"`
	Price         float64    `json:"price"`
	StockQuantity int        `json:"stock_quantity"`
	CreatedDate   *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
