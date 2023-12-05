package post

import (
	"time"
)

type Customer struct {
	ID           int        `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerName string     `json:"title" gorm:"size:25"`
	PhoneNumber  string     `json:"content" gorm:"type:text"`
	Email        string     `json:"category" gorm:"size:100"`
	CreatedDate  *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate  *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
