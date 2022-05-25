package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type networkPolicyService struct {
	networkPolicyRepo repository.NetworkPolicy
}

func (n networkPolicyService) Get(agent string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64) {
	return n.networkPolicyRepo.Get(agent, option)
}

func (n networkPolicyService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64) {
	return n.networkPolicyRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewNetworkPolicyService returns service.NetworkPolicy type service
func NewNetworkPolicyService(networkPolicyRepo repository.NetworkPolicy) service.NetworkPolicy {
	return &networkPolicyService{
		networkPolicyRepo: networkPolicyRepo,
	}
}
