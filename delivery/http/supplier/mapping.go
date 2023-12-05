package supplier

import (
	"kingkong-be/delivery/http/supplier/model"
	entity "kingkong-be/domain/supplier"
)

func mapRequestAddSupplier(req *model.Supplier, e *entity.Supplier) {
	e.SupplierID = req.SupplierID
	e.SupplierName = req.SupplierName
	e.PhoneNumber = req.PhoneNumber
	e.Email = req.Email
	e.ContactPerson = req.ContactPerson
}
