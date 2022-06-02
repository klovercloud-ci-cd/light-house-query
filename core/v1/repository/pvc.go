package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// PersistentVolumeClaim Repository operations.
type PersistentVolumeClaim interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64)
}
