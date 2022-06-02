package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// DaemonSet Repository operations.
type DaemonSet interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
}
