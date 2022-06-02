package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Namespace Repository operations.
type Namespace interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Namespace, int64)
}
