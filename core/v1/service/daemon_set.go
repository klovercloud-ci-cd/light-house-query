package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// DaemonSet business operations.
type DaemonSet interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64)
}
