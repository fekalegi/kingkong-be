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
	GetList(limit, offset int) ([]Transaction, int64, error)
	Get(id int) (*Transaction, error)
	Update(id int, req *Transaction) error
	Delete(id int) error
	AddTransactionPart(req *TransactionPart) error
	GetListPart(limit, offset int) ([]TransactionPart, int64, error)
	GetPart(id int) (*TransactionPart, error)
	UpdatePart(id int, req *TransactionPart) error
	DeletePart(id int) error
}

func (r *repository) AddTransaction(req *Transaction) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int) ([]Transaction, int64, error) {
	var transactions []Transaction
	var count int64

	query := r.db.Model(&transactions)

	err := query.Offset(offset).Limit(limit).Find(&transactions).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []Transaction{}, 0, err
	}

	return transactions, count, nil
}

func (r *repository) Get(id int) (*Transaction, error) {
	transaction := new(Transaction)

	if err := r.db.First(transaction, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) Update(id int, req *Transaction) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Transaction)
	return r.db.Where("id = ?", id).Delete(p).Error
}

func (r *repository) AddTransactionPart(req *TransactionPart) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetListPart(limit, offset int) ([]TransactionPart, int64, error) {
	var transactions []TransactionPart
	var count int64

	query := r.db.Model(&transactions)

	err := query.Offset(offset).Limit(limit).Find(&transactions).
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
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) DeletePart(id int) error {
	p := new(TransactionPart)
	return r.db.Where("id = ?", id).Delete(p).Error
}
