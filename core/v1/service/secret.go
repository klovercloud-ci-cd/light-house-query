package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Secret business operations.
type Secret interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Secret, int64)
}
