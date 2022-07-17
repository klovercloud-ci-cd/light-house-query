package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// PersistentVolumeClaim Repository operations.
type PersistentVolumeClaim interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64)
	GetById(id, agent, processId string) v1.PersistentVolumeClaim
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.PersistentVolumeClaim
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64)
}
