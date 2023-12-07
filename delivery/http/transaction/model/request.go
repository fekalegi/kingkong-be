package model

import (
	"time"
)

type Transaction struct {
	TransactionID         int               `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                int               `json:"user_id"`
	CustomerID            int               `json:"customer_id"`
	TransactionType       string            `json:"transaction_type"`
	TransactionDate       *time.Time        `json:"transaction_date"`
	TotalPrice            float64           `json:"total_price"`
	AdditionalInformation string            `json:"additional_information"`
	TransactionParts      []TransactionPart `json:"transaction_parts"`
}

type TransactionPart struct {
	TransactionPartID int        `json:"transaction_part_id"`
	TransactionID     int        `json:"transaction_id"`
	PartID            int        `json:"part_id"`
	Quantity          int        `json:"quantity"`
	Price             float64    `json:"price"`
	CreatedDate       *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}
