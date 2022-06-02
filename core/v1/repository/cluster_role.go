package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ClusterRole Repository operations.
type ClusterRole interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64)
	GetByProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRole, int64)
}
