package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// DaemonSet business operations.
type DaemonSet interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
	GetById(id, agent, processId string) v1.DaemonSet
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.DaemonSet
}
