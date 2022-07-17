package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ReplicaSet business operations.
type ReplicaSet interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64)
	GetById(id, agent, processId string) v1.ReplicaSet
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ReplicaSet
}
