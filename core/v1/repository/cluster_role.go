package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ClusterRole Repository operations.
type ClusterRole interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64)
	GetById(id, agent, processId string) v1.ClusterRole
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ClusterRole
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64)
}
