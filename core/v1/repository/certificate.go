package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Certificate Repository operations.
type Certificate interface {
	Get(agent string, option v1.ResourceQueryOption) ([]v1.Certificate, int64)
}
