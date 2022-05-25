package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type statefulSetService struct {
	statefulSetRepo repository.StatefulSet
}

func (s statefulSetService) Get(agent string, option v1.ResourceQueryOption) ([]v1.StatefulSet, int64) {
	return s.statefulSetRepo.Get(agent, option)
}

func (s statefulSetService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.StatefulSet, int64) {
	return s.statefulSetRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewStatefulSetService returns service.StatefulSet type service
func NewStatefulSetService(statefulSetRepo repository.StatefulSet) service.StatefulSet {
	return &statefulSetService{
		statefulSetRepo: statefulSetRepo,
	}
}
