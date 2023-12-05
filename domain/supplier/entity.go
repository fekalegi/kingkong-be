package supplier

import (
	"time"
)

type Supplier struct {
	SupplierID    int        `json:"supplier_id" gorm:"primaryKey;autoIncrement"`
	SupplierName  string     `json:"supplier_name"`
	PhoneNumber   string     `json:"phone_number"`
	Email         string     `json:"email"`
	ContactPerson string     `json:"contact_person"`
	CreatedDate   *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
