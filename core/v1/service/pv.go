package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// PersistentVolume business operations.
type PersistentVolume interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolume, int64)
	GetById(id, agent, processId string) v1.PersistentVolume
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.PersistentVolume
}
