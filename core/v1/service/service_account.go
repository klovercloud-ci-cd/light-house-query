package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ServiceAccount business operations.
type ServiceAccount interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64)
}
