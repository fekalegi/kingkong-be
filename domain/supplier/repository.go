package supplier

import (
	"errors"
	"gorm.io/gorm"
	"kingkong-be/common"
)

type repository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_supplier_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddSupplier(req *Supplier) error
	GetList(limit, offset int) ([]Supplier, int64, error)
	Get(id int) (*Supplier, error)
	Update(id int, req *Supplier) error
	Delete(id int) error
}

func (r *repository) AddSupplier(req *Supplier) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int) ([]Supplier, int64, error) {
	var suppliers []Supplier
	var count int64

	query := r.db.Model(&suppliers)

	err := query.Offset(offset).Limit(limit).Find(&suppliers).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []Supplier{}, 0, err
	}

	return suppliers, count, nil
}

func (r *repository) Get(id int) (*Supplier, error) {
	supplier := new(Supplier)

	if err := r.db.First(supplier, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (r *repository) Update(id int, req *Supplier) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Supplier)
	return r.db.Where("id = ?", id).Delete(p).Error
}
