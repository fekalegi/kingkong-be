package price_changes_log

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kingkong-be/common"
	"kingkong-be/helper"
)

type repository struct {
	db *gorm.DB
}

func NewPriceChangesLogRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_price_changes_log_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	Add(req *PriceChangesLog) error
	GetList(req *GetListRequest) ([]PriceChangesLog, int64, error)
	Get(id int) (*PriceChangesLog, error)
}

func (r *repository) Add(req *PriceChangesLog) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(req *GetListRequest) ([]PriceChangesLog, int64, error) {
	var priceChangesLogs []PriceChangesLog
	var count int64

	query := r.db.Model(&priceChangesLogs).Debug()

	if req.Search != "" {
		querySearch := fmt.Sprintf("parts.part_name LIKE '%%%[1]s%%' OR suppliers.supplier_name LIKE '%%%[1]s%%'", req.Search)
		query = query.Where(querySearch)
	}

	if req.StartDate != nil {
		query = query.Where("price_changes_logs.created_date >= ?", req.StartDate.In(helper.LoadLocationJakarta()))
	}

	if req.EndDate != nil {
		query = query.Where("price_changes_logs.created_date <= ?", req.EndDate.In(helper.LoadLocationJakarta()))
	}

	err := query.Offset(req.Offset).Limit(req.Limit).
		Select("id, parts.part_name, suppliers.supplier_name, price_before, price_after, price_changes_logs.created_date").
		Joins("LEFT JOIN parts ON parts.part_id = price_changes_logs.part_id").
		Joins("LEFT JOIN suppliers ON suppliers.supplier_id = parts.supplier_id").Order("price_changes_logs.created_date DESC").Find(&priceChangesLogs).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []PriceChangesLog{}, 0, err
	}

	return priceChangesLogs, count, nil
}

func (r *repository) Get(id int) (*PriceChangesLog, error) {
	priceChangesLog := new(PriceChangesLog)

	if err := r.db.First(priceChangesLog, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return priceChangesLog, nil
}

func (r *repository) Update(id int, req *PriceChangesLog) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(PriceChangesLog)
	return r.db.Where("id = ?", id).Delete(p).Error
}
