package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type serviceService struct {
	serviceRepo repository.Service
}

func (s serviceService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Service, int64) {
	return s.serviceRepo.Get(agent, option)
}

// NewServiceService returns service.Service type service
func NewServiceService(serviceRepo repository.Service) service.Service {
	return &serviceService{
		serviceRepo: serviceRepo,
	}
}