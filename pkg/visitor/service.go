package visitor

type service struct {
	repo Repo
}

type Service interface {
	Create(ip string) error
	List() ([]Visitor, error)
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Create(ip string) error {
	return s.repo.Create(&Visitor{IP: ip})
}

func (s *service) List() ([]Visitor, error) {
	return s.repo.List()
}
