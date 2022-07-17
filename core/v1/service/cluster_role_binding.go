package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ClusterRoleBinding business operations.
type ClusterRoleBinding interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64)
	GetById(id, agent, processId string) v1.ClusterRoleBinding
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.ClusterRoleBinding
}
