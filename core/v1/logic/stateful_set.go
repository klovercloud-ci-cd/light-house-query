package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type statefulSetService struct {
	statefulSetRepo repository.StatefulSet
}

func (s statefulSetService) GetById(id, agent, processId string) v1.StatefulSet {
	return s.statefulSetRepo.GetById(id, agent, processId)
}

func (s statefulSetService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.StatefulSet, int64) {
	if ownerReference == "" {
		return s.statefulSetRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return s.statefulSetRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (s statefulSetService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.StatefulSet {
	return s.statefulSetRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewStatefulSetService returns service.StatefulSet type service
func NewStatefulSetService(statefulSetRepo repository.StatefulSet) service.StatefulSet {
	return &statefulSetService{
		statefulSetRepo: statefulSetRepo,
	}
}
