package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type clusterRoleBindingService struct {
	clusterRoleBindingRepo repository.ClusterRoleBinding
}

func (c clusterRoleBindingService) Get(agent string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64) {
	return c.clusterRoleBindingRepo.Get(agent, option)
}

func (c clusterRoleBindingService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64) {
	return c.clusterRoleBindingRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewClusterRoleBindingService returns service.ClusterRoleBinding type service
func NewClusterRoleBindingService(clusterRoleBindingRepo repository.ClusterRoleBinding) service.ClusterRoleBinding {
	return &clusterRoleBindingService{
		clusterRoleBindingRepo: clusterRoleBindingRepo,
	}
}
