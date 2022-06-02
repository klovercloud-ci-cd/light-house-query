package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ReplicaSet Repository operations.
type ReplicaSet interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64)
}
