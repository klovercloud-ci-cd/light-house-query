package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Namespace business operations.
type Namespace interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Namespace, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Namespace, int64)
}
