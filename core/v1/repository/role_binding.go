package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// RoleBinding Repository operations.
type RoleBinding interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
	GetById(id, agent, processId string) v1.RoleBinding
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.RoleBinding
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
}
