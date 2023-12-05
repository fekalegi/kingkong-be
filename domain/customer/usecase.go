package post

import "sharing-vision-2021/common/helper"

type postImplementation struct {
	repo Repository
}

func NewPostImplementation(repo Repository) Service {
	return &postImplementation{
		repo: repo,
	}
}

type Service interface {
	AddPost(post *Post) error
	GetList(limit, offset int, status string) ([]Post, int64, error)
	Get(id int) (*Post, error)
	Update(id int, req *Post) error
	Delete(id int) error
}

func (u *postImplementation) AddPost(req *Post) error {
	return u.repo.AddPost(req)
}

func (u *postImplementation) GetList(limit, offset int, status string) ([]Post, int64, error) {
	return u.repo.GetList(limit, offset, helper.Capitalize(status))
}

func (u *postImplementation) Get(id int) (*Post, error) {
	return u.repo.Get(id)
}

func (u *postImplementation) Update(id int, req *Post) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *postImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
