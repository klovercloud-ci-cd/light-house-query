package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ServiceAccount business operations.
type ServiceAccount interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64)
	GetById(id, agent, processId string) v1.ServiceAccount
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ServiceAccount
}
