package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ConfigMap business operations.
type ConfigMap interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64)
	GetByProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.ConfigMap, int64)
}
