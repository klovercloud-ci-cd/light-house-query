package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// RoleBinding Repository operations.
type RoleBinding interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
}
