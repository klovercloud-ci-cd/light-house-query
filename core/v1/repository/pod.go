package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Pod Repository operations.
type Pod interface {
	Get(agent string) ([]v1.Pod, int64)
}
