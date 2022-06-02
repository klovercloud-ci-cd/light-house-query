package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// PersistentVolume business operations.
type PersistentVolume interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolume, int64)
}
