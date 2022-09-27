package domain

type DomainRepository interface {
	Example() (string, error)
}

type Service struct {
	logger     Logger
	repository DomainRepository
}

func NewService(repository DomainRepository, logger Logger) *Service {
	if repository == nil || logger == nil {
		return nil
	}
	return &Service{
		logger:     logger,
		repository: repository,
	}
}
func (s *Service) Example() (string, error) {
	return s.repository.Example()
}
