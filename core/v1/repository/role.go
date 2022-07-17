package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Role Repository operations.
type Role interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64)
	GetById(id, agent, processId string) v1.Role
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Role
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64)
}
