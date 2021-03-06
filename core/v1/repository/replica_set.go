package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ReplicaSet Repository operations.
type ReplicaSet interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64)
	GetById(id, agent, processId string) v1.ReplicaSet
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ReplicaSet
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64)
}
