package supplier

type supplierImplementation struct {
	repo Repository
}

func NewSupplierImplementation(repo Repository) Service {
	return &supplierImplementation{
		repo: repo,
	}
}

type Service interface {
	AddSupplier(supplier *Supplier) error
	GetList(limit, offset int) ([]Supplier, int64, error)
	Get(id int) (*Supplier, error)
	Update(id int, req *Supplier) error
	Delete(id int) error
}

func (u *supplierImplementation) AddSupplier(req *Supplier) error {
	return u.repo.AddSupplier(req)
}

func (u *supplierImplementation) GetList(limit, offset int) ([]Supplier, int64, error) {
	return u.repo.GetList(limit, offset)
}

func (u *supplierImplementation) Get(id int) (*Supplier, error) {
	return u.repo.Get(id)
}

func (u *supplierImplementation) Update(id int, req *Supplier) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *supplierImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
