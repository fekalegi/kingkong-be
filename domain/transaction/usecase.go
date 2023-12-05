package transaction

type transactionImplementation struct {
	repo Repository
}

func NewTransactionImplementation(repo Repository) Service {
	return &transactionImplementation{
		repo: repo,
	}
}

type Service interface {
	AddTransaction(transaction *Transaction) error
	GetList(limit, offset int) ([]Transaction, int64, error)
	Get(id int) (*Transaction, error)
	Update(id int, req *Transaction) error
	Delete(id int) error
}

func (u *transactionImplementation) AddTransaction(req *Transaction) error {
	return u.repo.AddTransaction(req)
}

func (u *transactionImplementation) GetList(limit, offset int) ([]Transaction, int64, error) {
	return u.repo.GetList(limit, offset)
}

func (u *transactionImplementation) Get(id int) (*Transaction, error) {
	return u.repo.Get(id)
}

func (u *transactionImplementation) Update(id int, req *Transaction) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return err
	}

	return u.repo.Update(id, req)
}

func (u *transactionImplementation) Delete(id int) error {
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
