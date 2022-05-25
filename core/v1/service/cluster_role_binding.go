package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// ClusterRoleBinding business operations.
type ClusterRoleBinding interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.ClusterRoleBinding, int64)
}
