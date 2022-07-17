package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// NetworkPolicy business operations.
type NetworkPolicy interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
	GetById(id, agent, processId string) v1.NetworkPolicy
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.NetworkPolicy
}
