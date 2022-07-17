package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// RoleBinding business operations.
type RoleBinding interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64)
	GetById(id, agent, processId string) v1.RoleBinding
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.RoleBinding
}
