package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Service business operations.
type Service interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Service, int64)
	GetById(id, agent, processId string) v1.Service
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Service
}
