package post

import (
	"time"
)

type Status string

const (
	Publish Status = "Publish"
	Draft   Status = "Draft"
	Thrash  Status = "Thrash"
)

type Post struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" gorm:"size:200"`
	Content     string     `json:"content" gorm:"type:text"`
	Category    string     `json:"category" gorm:"size:100"`
	CreatedDate *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
	Status      Status     `json:"status" gorm:"size:100;type:enum('Publish','Draft','Thrash')"`
}
