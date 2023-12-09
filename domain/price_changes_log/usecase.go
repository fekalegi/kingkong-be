package price_changes_log

type PriceChangesLogImplementation struct {
	repo Repository
}

func NewPriceChangesLogImplementation(repo Repository) Service {
	return &PriceChangesLogImplementation{
		repo: repo,
	}
}

type Service interface {
	GetList(req *GetListRequest) ([]PriceChangesLog, int64, error)
	Get(id int) (*PriceChangesLog, error)
}

func (u *PriceChangesLogImplementation) GetList(req *GetListRequest) ([]PriceChangesLog, int64, error) {
	return u.repo.GetList(req)
}

func (u *PriceChangesLogImplementation) Get(id int) (*PriceChangesLog, error) {
	return u.repo.Get(id)
}
