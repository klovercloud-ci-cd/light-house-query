package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type serviceAccountService struct {
	serviceAccountRepo repository.ServiceAccount
}

func (s serviceAccountService) GetById(id, agent, processId string) v1.ServiceAccount {
	return s.serviceAccountRepo.GetById(id, agent, processId)
}

func (s serviceAccountService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64) {
	if ownerReference == "" {
		return s.serviceAccountRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return s.serviceAccountRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (s serviceAccountService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ServiceAccount {
	return s.serviceAccountRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewServiceAccountService returns service.ServiceAccount type service
func NewServiceAccountService(serviceAccountRepo repository.ServiceAccount) service.ServiceAccount {
	return &serviceAccountService{
		serviceAccountRepo: serviceAccountRepo,
	}
}
