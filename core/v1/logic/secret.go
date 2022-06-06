package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type secretService struct {
	secretRepo repository.Secret
}

func (s secretService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Secret, int64) {
	if ownerReference == "" {
		return s.secretRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return s.secretRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

// NewSecretService returns service.Secret type service
func NewSecretService(secretRepo repository.Secret) service.Secret {
	return &secretService{
		secretRepo: secretRepo,
	}
}
