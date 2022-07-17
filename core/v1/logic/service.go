package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type serviceService struct {
	serviceRepo repository.Service
}

func (s serviceService) GetById(id, agent, processId string) v1.Service {
	return s.serviceRepo.GetById(id, agent, processId)
}

func (s serviceService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Service, int64) {
	if ownerReference == "" {
		return s.serviceRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return s.serviceRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (s serviceService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Service {
	return s.serviceRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewServiceService returns service.Service type service
func NewServiceService(serviceRepo repository.Service) service.Service {
	return &serviceService{
		serviceRepo: serviceRepo,
	}
}
