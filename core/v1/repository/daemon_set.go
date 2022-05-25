package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// DaemonSet Repository operations.
type DaemonSet interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
}
