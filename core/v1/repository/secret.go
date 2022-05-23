package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Secret Repository operations.
type Secret interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Secret, int64)
}
