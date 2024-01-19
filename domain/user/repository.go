package user

import (
	"errors"
	"gorm.io/gorm"
	"kingkong-be/common"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_user_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddUser(req *User) error
	GetList(limit, offset int) ([]User, int64, error)
	Get(id int) (*User, error)
	Update(id int, req *User) error
	Delete(id int) error
	GetByUsername(username string) (*User, error)
}

func (r *repository) AddUser(req *User) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int) ([]User, int64, error) {
	var users []User
	var count int64

	query := r.db.Model(&users)

	err := query.Offset(offset).Limit(limit).Find(&users).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []User{}, 0, err
	}

	return users, count, nil
}

func (r *repository) Get(id int) (*User, error) {
	user := new(User)

	if err := r.db.First(user, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(id int, req *User) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(User)
	return r.db.Where("id = ?", id).Delete(p).Error
}

func (r *repository) GetByUsername(username string) (*User, error) {
	user := new(User)

	if err := r.db.Where("username = ?", username).First(user).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
