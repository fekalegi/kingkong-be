package price_changes_log

import (
	"time"
)

type PriceChangesLog struct {
	ID            int        `json:"id"`
	PartID        int        `json:"part_id"`
	PartName      string     `json:"part_name" gorm:"->"`
	SupplierName  string     `json:"supplier_name" gorm:"->"`
	TransactionID int        `json:"transaction_id"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	CreatedDate   *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type GetListRequest struct {
	Limit     int        `json:"-" gorm:"-"`
	Offset    int        `json:"-" gorm:"-"`
	Search    string     `json:"-" gorm:"-"`
	StartDate *time.Time `json:"-" gorm:"-"`
	EndDate   *time.Time `json:"-" gorm:"-"`
}
