package post

import (
	"errors"
	"gorm.io/gorm"
	"sharing-vision-2021/common"
)

type repository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

//go:generate mockgen -destination=../../mocks/repository/mock_post_repository.go -package=mock_repository -source=repository.go
type Repository interface {
	AddPost(req *Post) error
	GetList(limit, offset int, status string) ([]Post, int64, error)
	Get(id int) (*Post, error)
	Update(id int, req *Post) error
	Delete(id int) error
}

func (r *repository) AddPost(req *Post) error {
	return r.db.Create(&req).Error
}

func (r *repository) GetList(limit, offset int, status string) ([]Post, int64, error) {
	var posts []Post
	var count int64

	query := r.db.Model(&posts)

	if status != "" {
		query.Where("status = ?", Status(status))
	}

	err := query.Offset(offset).Limit(limit).Find(&posts).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []Post{}, 0, err
	}

	return posts, count, nil
}

func (r *repository) Get(id int) (*Post, error) {
	post := new(Post)

	if err := r.db.First(post, id).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *repository) Update(id int, req *Post) error {
	return r.db.Model(req).Where("id = ?", id).Updates(&req).Error
}

func (r *repository) Delete(id int) error {
	p := new(Post)
	return r.db.Where("id = ?", id).Delete(p).Error
}
