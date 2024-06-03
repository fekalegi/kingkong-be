package transaction

import (
	"errors"
	"gorm.io/gorm"
	"kingkong-be/common"
)

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_transaction_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddTransaction(req *Transaction) error
	GetList(limit, offset int, trxType string) ([]TransactionReport, int64, error)
	Get(id int) (*TransactionReport, error)
	Update(id int, req *Transaction) error
	Delete(id int) error
	AddBatchTransactionPart(req []TransactionPart) error
	GetListPart(limit, offset, transactionID int) ([]TransactionPart, int64, error)
	GetPart(id int) (*TransactionPart, error)
	UpdatePart(id int, req *TransactionPart) error
	DeletePart(id int) error
	DeletePartsByTransactionID(id int) error
	GetSum7DaysBefore(status string) ([]WeeklyChart, error)
	GetSumMonthly(status string) ([]MonthlyChart, error)
}

func (r *repository) AddTransaction(req *Transaction) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int, trxType string) ([]TransactionReport, int64, error) {
	var transactions []TransactionReport
	var count int64

	query := r.db.Model(&transactions)
	if trxType != "" {
		query.Where("transactions.transaction_type = ?", trxType)
	}
	err := query.Offset(offset).Limit(limit).
		Select("transactions.transaction_id, transactions.user_id, users.username, transactions.customer_id, customers.customer_name, transaction_type, transaction_date, total_price, additional_information").
		Joins("LEFT JOIN users ON transactions.user_id = users.user_id").
		Joins("LEFT JOIN customers ON transactions.customer_id = customers.customer_id").
		Find(&transactions).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []TransactionReport{}, 0, err
	}

	return transactions, count, nil
}

func (r *repository) Get(id int) (*TransactionReport, error) {
	transaction := new(TransactionReport)

	if err := r.db.
		Select("transactions.transaction_id, transactions.user_id, users.username, transactions.customer_id, customers.customer_name, transaction_type, transaction_date, total_price, additional_information").
		Joins("LEFT JOIN users ON transactions.user_id = users.user_id").
		Joins("LEFT JOIN customers ON transactions.customer_id = customers.customer_id").
		First(transaction, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) Update(id int, req *Transaction) error {
	return r.db.Model(req).Where("transaction_id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Transaction)
	return r.db.Where("transaction_id = ?", id).Delete(p).Error
}

func (r *repository) AddBatchTransactionPart(req []TransactionPart) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetListPart(limit, offset, transactionID int) ([]TransactionPart, int64, error) {
	var transactions []TransactionPart
	var count int64

	query := r.db.Model(&transactions)

	err := query.Offset(offset).Limit(limit).
		Select("transaction_part_id, transaction_id, parts.part_id, parts.part_name, quantity, transaction_parts.price, "+
			"transaction_parts.created_date, transaction_parts.updated_date").
		Joins("LEFT JOIN parts ON parts.part_id = transaction_parts.part_id").
		Where("transaction_id = ?", transactionID).Find(&transactions).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []TransactionPart{}, 0, err
	}

	return transactions, count, nil
}

func (r *repository) GetPart(id int) (*TransactionPart, error) {
	transaction := new(TransactionPart)

	if err := r.db.First(transaction, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) UpdatePart(id int, req *TransactionPart) error {
	return r.db.Model(req).Where("transaction_part_id = ?", id).Updates(&req).Error
}

func (r *repository) DeletePart(id int) error {
	p := new(TransactionPart)
	return r.db.Where("transaction_part_id = ?", id).Delete(p).Error
}

func (r *repository) DeletePartsByTransactionID(id int) error {
	p := new(TransactionPart)
	return r.db.Where("transaction_id = ?", id).Delete(p).Error
}

func (r *repository) GetSum7DaysBefore(status string) ([]WeeklyChart, error) {
	var res []WeeklyChart

	if err := r.db.Raw(
		"SELECT DAYOFWEEK(transaction_date) as day_of_week,  SUM(total_price) as sum "+
			"FROM transactions WHERE transaction_type = ? AND transaction_date >= CURRENT_DATE - INTERVAL 7 DAY "+
			"GROUP BY day_of_week "+
			"ORDER BY  day_of_week;", status).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetSumMonthly(status string) ([]MonthlyChart, error) {
	var res []MonthlyChart

	if err := r.db.Raw(
		"SELECT EXTRACT(YEAR_MONTH FROM transaction_date) AS month_year ,"+
			"EXTRACT(MONTH FROM transaction_date)      AS month,"+
			"SUM(total_price)                          AS sum "+
			"FROM transactions "+
			"WHERE transaction_type = ? "+
			"AND transaction_date >= CURRENT_DATE - INTERVAL 12 MONTH "+
			"GROUP BY month_year "+
			"ORDER BY month_year;", status).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
