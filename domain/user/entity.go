package user

import (
	"time"
)

type User struct {
	UserID      int        `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Role        int        `json:"role"`
	CreatedDate *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
