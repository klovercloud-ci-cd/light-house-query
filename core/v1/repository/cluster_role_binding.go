package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ClusterRoleBinding Repository operations.
type ClusterRoleBinding interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64)
	GetById(id, agent, processId string) v1.ClusterRoleBinding
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ClusterRoleBinding
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64)
}
