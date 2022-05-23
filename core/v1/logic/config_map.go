package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type configMapService struct {
	configMapRepo repository.ConfigMap
}

func (c configMapService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64) {
	return c.configMapRepo.Get(agent, option)
}

// NewConfigMapService returns service.ConfigMap type service
func NewConfigMapService(configMapRepo repository.ConfigMap) service.ConfigMap {
	return &configMapService{
		configMapRepo: configMapRepo,
	}
}
