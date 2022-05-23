package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Node Repository operations.
type Node interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Node, int64)
}
