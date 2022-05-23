package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type replicaSetService struct {
	replicaSetServiceRepo repository.ReplicaSet
}

func (r replicaSetService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64) {
	return r.replicaSetServiceRepo.Get(agent, option)
}

// NewReplicaSetService returns service.ReplicaSetService type service
func NewReplicaSetService(replicaSetRepo repository.ReplicaSet) service.ReplicaSet {
	return &replicaSetService{
		replicaSetServiceRepo: replicaSetRepo,
	}
}
