package user

type userImplementation struct {
	repo Repository
}

func NewUserImplementation(repo Repository) Service {
	return &userImplementation{
		repo: repo,
	}
}

type Service interface {
	AddUser(user *User) error
	GetList(limit, offset int) ([]User, int64, error)
	Get(id int) (*User, error)
	Update(id int, req *User) error
	Delete(id int) error
}

func (u *userImplementation) AddUser(req *User) error {
	return u.repo.AddUser(req)
}

func (u *userImplementation) GetList(limit, offset int) ([]User, int64, error) {
	return u.repo.GetList(limit, offset)
}

func (u *userImplementation) Get(id int) (*User, error) {
	return u.repo.Get(id)
}

func (u *userImplementation) Update(id int, req *User) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *userImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
