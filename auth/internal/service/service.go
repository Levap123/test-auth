package service

type Service struct {
	UserService
}

func NewService() *Service {
	return &Service{}
}

type UserService interface {
	Create(email, password string) error
}
