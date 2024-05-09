package account

type Service interface {
	GetAccounts(limit int64) ([]Account, error)
	GetAccountByID(id int64) (*Account, error)
	GetAccountByUsername(username string) (*Account, error)
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repo Repository) *ServiceImpl {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) GetAccounts(limit int64) ([]Account, error) {
	return s.repo.GetAccounts(limit)
}

func (s *ServiceImpl) GetAccountByID(id int64) (*Account, error) {
	return s.repo.GetAccountByID(id)
}

func (s *ServiceImpl) GetAccountByUsername(username string) (*Account, error) {
	return s.repo.GetAccountByUsername(username)
}
