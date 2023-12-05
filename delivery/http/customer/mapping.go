package customer

import (
	"kingkong-be/delivery/http/customer/model"
	entity "kingkong-be/domain/customer"
)

func mapRequestAddCustomer(req *model.Customer, e *entity.Customer) {
	e.CustomerID = req.CustomerID
	e.CustomerName = req.CustomerName
	e.PhoneNumber = req.PhoneNumber
	e.Email = req.Email
}
