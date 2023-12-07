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
	AddBatchTransactionPart(transaction []TransactionPart) error
	GetListPart(limit, offset int) ([]TransactionPart, int64, error)
	GetPart(id int) (*TransactionPart, error)
	UpdateBatchPart(id int, req []TransactionPart) error
	DeletePart(id int) error
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

func (u *transactionImplementation) AddBatchTransactionPart(req []TransactionPart) error {
	return u.repo.AddBatchTransactionPart(req)
}

func (u *transactionImplementation) GetListPart(limit, offset int) ([]TransactionPart, int64, error) {
	return u.repo.GetListPart(limit, offset)
}

func (u *transactionImplementation) GetPart(id int) (*TransactionPart, error) {
	return u.repo.GetPart(id)
}

func (u *transactionImplementation) UpdateBatchPart(id int, req []TransactionPart) error {
	err := u.repo.DeletePartsByTransactionID(id)
	if err != nil {
		return err
	}

	return u.repo.AddBatchTransactionPart(req)
}

func (u *transactionImplementation) DeletePart(id int) error {
	if _, err := u.repo.GetPart(id); err != nil {
		return err
	}

	return u.repo.DeletePart(id)
}
