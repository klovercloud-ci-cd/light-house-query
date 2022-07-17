package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type networkPolicyService struct {
	networkPolicyRepo repository.NetworkPolicy
}

func (n networkPolicyService) GetById(id, agent, processId string) v1.NetworkPolicy {
	return n.networkPolicyRepo.GetById(id, agent, processId)
}

func (n networkPolicyService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64) {
	if ownerReference == "" {
		return n.networkPolicyRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return n.networkPolicyRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (n networkPolicyService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.NetworkPolicy {
	return n.networkPolicyRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewNetworkPolicyService returns service.NetworkPolicy type service
func NewNetworkPolicyService(networkPolicyRepo repository.NetworkPolicy) service.NetworkPolicy {
	return &networkPolicyService{
		networkPolicyRepo: networkPolicyRepo,
	}
}
