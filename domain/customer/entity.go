package customer

import (
	"time"
)

type Customer struct {
	CustomerID   int        `json:"customer_id" gorm:"primaryKey;autoIncrement"`
	CustomerName string     `json:"title"`
	PhoneNumber  string     `json:"content"`
	Email        string     `json:"category"`
	CreatedDate  *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate  *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
