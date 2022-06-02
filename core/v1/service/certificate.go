package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Certificate business operations.
type Certificate interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
	GetByProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
}
