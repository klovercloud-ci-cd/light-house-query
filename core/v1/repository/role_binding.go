package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// RoleBinding Repository operations.
type RoleBinding interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
}
