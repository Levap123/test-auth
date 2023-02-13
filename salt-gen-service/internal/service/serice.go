package service

type Service struct {
	SaltService
}

func NewService() *Service {
	return &Service{
		SaltService: NewSalt(),
	}
}

type SaltService interface {
	GenerateSalt() string
}
