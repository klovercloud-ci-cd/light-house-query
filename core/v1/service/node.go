package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Node business operations.
type Node interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Node, int64)
	GetById(id, agent, processId string) v1.Node
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Node
}
