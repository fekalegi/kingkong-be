package model

type User struct {
	UserID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=-1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
