package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// NetworkPolicy Repository operations.
type NetworkPolicy interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64)
}
