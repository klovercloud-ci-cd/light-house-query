package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Secret Repository operations.
type Secret interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Secret, int64)
	GetById(id, agent, processId string) v1.Secret
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Secret
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Secret, int64)
}
