package customer

import (
	"time"
)

type Customer struct {
	CustomerID   int        `json:"customer_id" gorm:"primaryKey;autoIncrement"`
	CustomerName string     `json:"customer_name"`
	PhoneNumber  string     `json:"phone_number"`
	Email        string     `json:"email"`
	CreatedDate  *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate  *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
