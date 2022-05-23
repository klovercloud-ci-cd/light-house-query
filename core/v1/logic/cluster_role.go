package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type clusterRoleService struct {
	clusterRoleRepo repository.ClusterRole
}

func (c clusterRoleService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64) {
	return c.clusterRoleRepo.Get(agent, option)
}

// NewClusterRoleService returns service.ClusterRole type service
func NewClusterRoleService(clusterRoleRepo service.ClusterRole) service.ClusterRole {
	return &clusterRoleService{
		clusterRoleRepo: clusterRoleRepo,
	}
}
