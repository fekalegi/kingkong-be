package model

type Post struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}
