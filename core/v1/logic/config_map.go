package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type configMapService struct {
	configMapRepo repository.ConfigMap
}

func (c configMapService) GetById(id, agent, processId string) v1.ConfigMap {
	return c.configMapRepo.GetById(id, agent, processId)
}

func (c configMapService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64) {
	if ownerReference == "" {
		return c.configMapRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return c.configMapRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (c configMapService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ConfigMap {
	return c.configMapRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewConfigMapService returns service.ConfigMap type service
func NewConfigMapService(configMapRepo repository.ConfigMap) service.ConfigMap {
	return &configMapService{
		configMapRepo: configMapRepo,
	}
}
