package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Certificate Repository operations.
type Certificate interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
	GetById(id, agent, processId string) v1.Certificate
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Certificate
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
}
