package part

import (
	"errors"
	"gorm.io/gorm"
	"kingkong-be/common"
)

type repository struct {
	db *gorm.DB
}

func NewPartRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_part_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddPart(req *Part) error
	GetList(limit, offset int) ([]Part, int64, error)
	Get(id int) (*Part, error)
	Update(id int, req *Part) error
	Delete(id int) error
}

func (r *repository) AddPart(req *Part) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int) ([]Part, int64, error) {
	var parts []Part
	var count int64

	query := r.db.Model(&parts)

	err := query.Offset(offset).Limit(limit).Find(&parts).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []Part{}, 0, err
	}

	return parts, count, nil
}

func (r *repository) Get(id int) (*Part, error) {
	part := new(Part)

	if err := r.db.First(part, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return part, nil
}

func (r *repository) Update(id int, req *Part) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Part)
	return r.db.Where("id = ?", id).Delete(p).Error
}
