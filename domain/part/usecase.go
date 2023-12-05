package part

type partImplementation struct {
	repo Repository
}

func NewPartImplementation(repo Repository) Service {
	return &partImplementation{
		repo: repo,
	}
}

type Service interface {
	AddPart(part *Part) error
	GetList(limit, offset int) ([]Part, int64, error)
	Get(id int) (*Part, error)
	Update(id int, req *Part) error
	Delete(id int) error
}

func (u *partImplementation) AddPart(req *Part) error {
	return u.repo.AddPart(req)
}

func (u *partImplementation) GetList(limit, offset int) ([]Part, int64, error) {
	return u.repo.GetList(limit, offset)
}

func (u *partImplementation) Get(id int) (*Part, error) {
	return u.repo.Get(id)
}

func (u *partImplementation) Update(id int, req *Part) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *partImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
