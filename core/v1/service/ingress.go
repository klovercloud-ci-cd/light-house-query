package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Ingress business operations.
type Ingress interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Ingress, int64)
	GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Ingress, int64)
}
