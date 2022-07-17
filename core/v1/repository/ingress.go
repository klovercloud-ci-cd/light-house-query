package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Ingress Repository operations.
type Ingress interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Ingress, int64)
	GetById(id, agent, processId string) v1.Ingress
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Ingress
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Ingress, int64)
}
