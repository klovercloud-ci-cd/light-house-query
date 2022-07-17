package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Namespace business operations.
type Namespace interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Namespace, int64)
	GetById(id, agent, processId string) v1.Namespace
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Namespace
}
