package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// DaemonSet Repository operations.
type DaemonSet interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
	GetById(id, agent, processId string) v1.DaemonSet
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.DaemonSet
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
}
