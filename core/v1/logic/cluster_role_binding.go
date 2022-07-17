package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type clusterRoleBindingService struct {
	clusterRoleBindingRepo repository.ClusterRoleBinding
}

func (c clusterRoleBindingService) GetById(id, agent, processId string) v1.ClusterRoleBinding {
	return c.clusterRoleBindingRepo.GetById(id, agent, processId)
}

func (c clusterRoleBindingService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64) {
	if ownerReference == "" {
		return c.clusterRoleBindingRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return c.clusterRoleBindingRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (c clusterRoleBindingService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ClusterRoleBinding {
	return c.clusterRoleBindingRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewClusterRoleBindingService returns service.ClusterRoleBinding type service
func NewClusterRoleBindingService(clusterRoleBindingRepo repository.ClusterRoleBinding) service.ClusterRoleBinding {
	return &clusterRoleBindingService{
		clusterRoleBindingRepo: clusterRoleBindingRepo,
	}
}
