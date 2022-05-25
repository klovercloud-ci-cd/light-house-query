package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ServiceAccount Repository operations.
type ServiceAccount interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64)
}
