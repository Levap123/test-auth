package transport

import (
	"github.com/Levap123/auth-service/internal/service"
	"github.com/Levap123/auth-service/internal/validator"
)

type Transport struct {
	service   *service.Service
	validator *validator.Validator
}

func NewTransport(service *service.Service, validator *validator.Validator) *Transport {
	return &Transport{
		service:   service,
		validator: validator,
	}
}
