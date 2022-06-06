package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Node Repository operations.
type Node interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Node, int64)
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Node, int64)
}
