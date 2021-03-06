package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Deployment business operations.
type Deployment interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64)
	GetById(id, agent, processId string) v1.Deployment
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Deployment
}
