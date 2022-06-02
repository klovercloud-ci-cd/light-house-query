package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ServiceAccount Repository operations.
type ServiceAccount interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64)
}
