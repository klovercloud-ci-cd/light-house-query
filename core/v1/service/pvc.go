package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// PersistentVolumeClaim business operations.
type PersistentVolumeClaim interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64)
}
