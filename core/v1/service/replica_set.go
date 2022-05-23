package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ReplicaSet business operations.
type ReplicaSet interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64)
}
