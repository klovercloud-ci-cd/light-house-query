package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Agent repository operations.
type Agent interface {
	Get(companyId string) []v1.Agent
}
