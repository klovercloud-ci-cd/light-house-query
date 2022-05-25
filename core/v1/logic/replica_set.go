package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type replicaSetService struct {
	replicaSetRepo repository.ReplicaSet
}

func (r replicaSetService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64) {
	return r.replicaSetRepo.Get(agent, option)
}

func (r replicaSetService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64) {
	return r.replicaSetRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewReplicaSetService returns service.ReplicaSet type service
func NewReplicaSetService(replicaSetRepo repository.ReplicaSet) service.ReplicaSet {
	return &replicaSetService{
		replicaSetRepo: replicaSetRepo,
	}
}
