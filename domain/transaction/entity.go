package transaction

import (
	"time"
)

type Transaction struct {
	TransactionID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                int        `json:"user_id"`
	CustomerID            int        `json:"customer_id"`
	TransactionType       string     `json:"transaction_type"`
	TransactionDate       *time.Time `json:"transaction_date"`
	TotalPrice            float64    `json:"total_price"`
	AdditionalInformation string     `json:"additional_information"`
	CreatedDate           *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate           *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type TransactionPart struct {
	TransactionPartID int        `json:"transaction_part_id" gorm:"primaryKey;autoIncrement"`
	TransactionID     int        `json:"transaction_id"`
	PartID            int        `json:"part_id"`
	PartName          string     `json:"part_name" gorm:"->"`
	Quantity          int        `json:"quantity"`
	Price             float64    `json:"price"`
	CreatedDate       *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type TransactionReport struct {
	TransactionID         int               `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	UserID                int               `json:"user_id"`
	Username              string            `json:"username"`
	CustomerID            int               `json:"customer_id"`
	CustomerName          string            `json:"customer_name"`
	TransactionType       string            `json:"transaction_type"`
	TransactionDate       *time.Time        `json:"transaction_date"`
	TotalPrice            float64           `json:"total_price"`
	AdditionalInformation string            `json:"additional_information"`
	TransactionParts      []TransactionPart `json:"transaction_parts" gorm:"foreignKey:transaction_id"`
	CreatedDate           *time.Time        `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate           *time.Time        `json:"updated_date" gorm:"autoUpdateTime"`
}

type RequestInsertTransaction struct {
	Transaction      Transaction       `json:"transaction"`
	TransactionParts []TransactionPart `json:"transaction_parts"`
}

func (tr *TransactionReport) TableName() string {
	return "transactions"
}
