package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// RoleBinding business operations.
type RoleBinding interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
}
