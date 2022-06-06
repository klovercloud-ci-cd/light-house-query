package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type replicaSetService struct {
	replicaSetRepo repository.ReplicaSet
}

func (r replicaSetService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64) {
	if ownerReference == "" {
		return r.replicaSetRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return r.replicaSetRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

// NewReplicaSetService returns service.ReplicaSet type service
func NewReplicaSetService(replicaSetRepo repository.ReplicaSet) service.ReplicaSet {
	return &replicaSetService{
		replicaSetRepo: replicaSetRepo,
	}
}
