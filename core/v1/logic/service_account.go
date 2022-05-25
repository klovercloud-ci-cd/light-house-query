package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type serviceAccountService struct {
	serviceAccountRepo repository.ServiceAccount
}

func (s serviceAccountService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64) {
	return s.serviceAccountRepo.Get(agent, option)
}
func (s serviceAccountService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64) {
	return s.serviceAccountRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewServiceAccountService returns service.ServiceAccount type service
func NewServiceAccountService(serviceAccountRepo repository.ServiceAccount) service.ServiceAccount {
	return &serviceAccountService{
		serviceAccountRepo: serviceAccountRepo,
	}
}
