package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Node business operations.
type Node interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Node, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Node, int64)
}
