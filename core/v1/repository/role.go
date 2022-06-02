package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Role Repository operations.
type Role interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64)
}
