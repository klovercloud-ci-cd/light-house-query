package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ConfigMap Repository operations.
type ConfigMap interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64)
	GetById(id, agent, processId string) v1.ConfigMap
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ConfigMap
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64)
}
