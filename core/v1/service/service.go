package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Service business operations.
type Service interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Service, int64)
}
