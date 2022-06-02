package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Role business operations.
type Role interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64)
}
