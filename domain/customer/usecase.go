package customer

type customerImplementation struct {
	repo Repository
}

func NewCustomerImplementation(repo Repository) Service {
	return &customerImplementation{
		repo: repo,
	}
}

type Service interface {
	AddCustomer(customer *Customer) error
	GetList(limit, offset int) ([]Customer, int64, error)
	Get(id int) (*Customer, error)
	Update(id int, req *Customer) error
	Delete(id int) error
}

func (u *customerImplementation) AddCustomer(req *Customer) error {
	return u.repo.AddCustomer(req)
}

func (u *customerImplementation) GetList(limit, offset int) ([]Customer, int64, error) {
	return u.repo.GetList(limit, offset)
}

func (u *customerImplementation) Get(id int) (*Customer, error) {
	return u.repo.Get(id)
}

func (u *customerImplementation) Update(id int, req *Customer) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *customerImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
