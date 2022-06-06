package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type clusterRoleService struct {
	clusterRoleRepo repository.ClusterRole
}

func (c clusterRoleService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64) {
	if ownerReference == "" {
		return c.clusterRoleRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return c.clusterRoleRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

// NewClusterRoleService returns service.ClusterRole type service
func NewClusterRoleService(clusterRoleRepo repository.ClusterRole) service.ClusterRole {
	return &clusterRoleService{
		clusterRoleRepo: clusterRoleRepo,
	}
}
