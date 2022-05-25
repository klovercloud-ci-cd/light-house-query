package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Pod Repository operations.
type Pod interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Pod, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Pod, int64)
}
