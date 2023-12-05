package customer

import (
	"errors"
	"gorm.io/gorm"
	"kingkong-be/common"
)

type repository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_customer_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddCustomer(req *Customer) error
	GetList(limit, offset int) ([]Customer, int64, error)
	Get(id int) (*Customer, error)
	Update(id int, req *Customer) error
	Delete(id int) error
}

func (r *repository) AddCustomer(req *Customer) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int) ([]Customer, int64, error) {
	var customers []Customer
	var count int64

	query := r.db.Model(&customers)

	err := query.Offset(offset).Limit(limit).Find(&customers).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []Customer{}, 0, err
	}

	return customers, count, nil
}

func (r *repository) Get(id int) (*Customer, error) {
	customer := new(Customer)

	if err := r.db.First(customer, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *repository) Update(id int, req *Customer) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Customer)
	return r.db.Where("id = ?", id).Delete(p).Error
}
