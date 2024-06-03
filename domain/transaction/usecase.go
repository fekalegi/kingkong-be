package transaction

import (
	"kingkong-be/common"
	"kingkong-be/domain/part"
	priceLog "kingkong-be/domain/price_changes_log"
)

type transactionImplementation struct {
	repo     Repository
	partRepo part.Repository
	priceLog priceLog.Repository
}

func NewTransactionImplementation(repo Repository,
	partRepo part.Repository,
	priceLog priceLog.Repository) Service {
	return &transactionImplementation{
		repo:     repo,
		partRepo: partRepo,
		priceLog: priceLog,
	}
}

type Service interface {
	AddTransaction(transaction *RequestInsertTransaction) error
	GetList(limit, offset int, trxType string) ([]TransactionReport, int64, error)
	Get(id int) (*TransactionReport, error)
	Update(id int, req *Transaction) error
	Delete(id int) error
	AddBatchTransactionPart(transaction []TransactionPart) error
	GetPart(id int) (*TransactionPart, error)
	UpdateBatchPart(id int, req []TransactionPart) error
	DeletePart(id int) error
	GetSumSales7DaysBefore() ([]WeeklyChart, error)
	GetSumPurchase7DaysBefore() ([]WeeklyChart, error)
	GetSumSalesMonthly() ([]MonthlyChart, error)
	GetSumPurchaseMonthly() ([]MonthlyChart, error)
}

func (u *transactionImplementation) AddTransaction(req *RequestInsertTransaction) error {
	err := u.repo.AddTransaction(&req.Transaction)
	if err != nil {
		return err
	}

	id := req.Transaction.TransactionID

	totalPrice := 0.00
	for k, v := range req.TransactionParts {
		totalPrice += v.Price * float64(v.Quantity)
		quantity := v.Quantity
		pt, err := u.partRepo.Get(v.PartID)
		if err != nil {
			return err
		}

		if pt.Price != v.Price && req.Transaction.TransactionType == common.Purchase {
			priceChanges := &priceLog.PriceChangesLog{
				PartID:        v.PartID,
				TransactionID: id,
				PriceBefore:   pt.Price,
				PriceAfter:    v.Price,
			}

			err = u.priceLog.Add(priceChanges)
		}

		if err != nil {
			return err
		}

		if req.Transaction.TransactionType == common.Sales {
			quantity = -1 * v.Quantity
		}

		err = u.partRepo.UpdateStockByID(v.PartID, quantity)
		if err != nil {
			return err
		}

		err = u.partRepo.Update(v.PartID, &part.Part{
			Price: v.Price,
		})
		if err != nil {
			return err
		}

		err = u.partRepo.Update(v.PartID, &part.Part{
			Price: v.Price,
		})
		if err != nil {
			return err
		}

		req.TransactionParts[k].TransactionID = id
	}

	err = u.repo.AddBatchTransactionPart(req.TransactionParts)
	if err != nil {
		return err
	}

	req.Transaction.TotalPrice = totalPrice
	return u.repo.Update(id, &req.Transaction)
}

func (u *transactionImplementation) GetList(limit, offset int, trxType string) ([]TransactionReport, int64, error) {
	return u.repo.GetList(limit, offset, trxType)
}

func (u *transactionImplementation) Get(id int) (*TransactionReport, error) {
	trx, err := u.repo.Get(id)
	if err != nil {
		return nil, err
	}

	tps, _, err := u.repo.GetListPart(-1, 0, id)
	if err != nil {
		return nil, err
	}

	res := &TransactionReport{
		TransactionID:         trx.TransactionID,
		UserID:                trx.UserID,
		Username:              trx.Username,
		CustomerID:            trx.CustomerID,
		CustomerName:          trx.CustomerName,
		TransactionType:       trx.TransactionType,
		TransactionDate:       trx.TransactionDate,
		TotalPrice:            trx.TotalPrice,
		AdditionalInformation: trx.AdditionalInformation,
		TransactionParts:      tps,
	}

	return res, nil
}

func (u *transactionImplementation) Update(id int, req *Transaction) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *transactionImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}

func (u *transactionImplementation) AddBatchTransactionPart(req []TransactionPart) error {
	return u.repo.AddBatchTransactionPart(req)
}

func (u *transactionImplementation) GetPart(id int) (*TransactionPart, error) {
	return u.repo.GetPart(id)
}

func (u *transactionImplementation) UpdateBatchPart(id int, req []TransactionPart) error {
	err := u.repo.DeletePartsByTransactionID(id)
	if err != nil {
		return err
	}

	return u.repo.AddBatchTransactionPart(req)
}

func (u *transactionImplementation) DeletePart(id int) error {
	if _, err := u.repo.GetPart(id); err != nil {
		return err
	}

	return u.repo.DeletePart(id)
}

func (u *transactionImplementation) GetSumSales7DaysBefore() ([]WeeklyChart, error) {
	wcs, err := u.repo.GetSum7DaysBefore("sales")
	if err != nil {
		return nil, err
	}
	return generateEmptyWeeklyChart(wcs), err
}

func (u *transactionImplementation) GetSumPurchase7DaysBefore() ([]WeeklyChart, error) {
	wcs, err := u.repo.GetSum7DaysBefore("purchase")
	if err != nil {
		return nil, err
	}
	return generateEmptyWeeklyChart(wcs), err
}

func (u *transactionImplementation) GetSumSalesMonthly() ([]MonthlyChart, error) {
	wcs, err := u.repo.GetSumMonthly("sales")
	if err != nil {
		return nil, err
	}
	return generateEmptyMonthlyChart(wcs), err
}

func (u *transactionImplementation) GetSumPurchaseMonthly() ([]MonthlyChart, error) {
	wcs, err := u.repo.GetSumMonthly("purchase")
	if err != nil {
		return nil, err
	}
	return generateEmptyMonthlyChart(wcs), err
}
