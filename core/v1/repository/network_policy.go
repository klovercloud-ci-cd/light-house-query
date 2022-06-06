package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// NetworkPolicy Repository operations.
type NetworkPolicy interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
}
