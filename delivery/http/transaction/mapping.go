package transaction

import (
	"kingkong-be/delivery/http/transaction/model"
	entity "kingkong-be/domain/transaction"
)

func mapRequestAddTransaction(req *model.Transaction, e *entity.Transaction) {
	e.TransactionID = req.TransactionID
	e.UserID = req.UserID
	e.CustomerID = req.CustomerID
	e.TransactionType = req.TransactionType
	e.TransactionDate = req.TransactionDate
	e.TotalPrice = req.TotalPrice
	e.AdditionalInformation = req.AdditionalInformation
}

func mapRequestAddTransactionPart(req *model.TransactionPart, e *entity.TransactionPart) {
	e.TransactionPartID = req.TransactionPartID
	e.TransactionID = req.TransactionID
	e.PartID = req.PartID
	e.Quantity = req.Quantity
	e.Price = req.Price
}
