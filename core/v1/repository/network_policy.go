package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// NetworkPolicy Repository operations.
type NetworkPolicy interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
	GetById(id, agent, processId string) v1.NetworkPolicy
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.NetworkPolicy
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
}
