package model

type Supplier struct {
	SupplierID    int    `json:"supplier_id" gorm:"primaryKey;autoIncrement"`
	SupplierName  string `json:"supplier_name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	ContactPerson string `json:"contact_person"`
}

type List struct {
	Limit  int    `form:"limit" validate:"required,min=-1"`
	Offset int    `form:"offset" validate:"min=0"`
	Status string `form:"status" validate:"oneof='publish' 'draft' 'thrash' ''"`
}
