package model

type Customer struct {
	CustomerID   int    `json:"customer_id"`
	CustomerName string `json:"title"`
	PhoneNumber  string `json:"content"`
	Email        string `json:"category"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}
