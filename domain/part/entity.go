package part

import (
	"time"
)

type Part struct {
	PartID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	PartName      string     `json:"part_name"`
	Price         float64    `json:"price"`
	StockQuantity int        `json:"stock_quantity"`
	CreatedDate   *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
