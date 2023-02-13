package transport

import "github.com/Levap123/salt-gen/internal/service"

type Transport struct {
	service *service.Service
}

func NewTransport(service *service.Service) *Transport {
	return &Transport{
		service: service,
	}
}
