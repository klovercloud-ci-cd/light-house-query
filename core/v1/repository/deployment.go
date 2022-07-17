package repository

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Deployment Repository operations.
type Deployment interface {
	GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64)
	GetById(id, agent, processId string) v1.Deployment
	GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Deployment
	GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64)
	CountDeploymentsByCompanyIdAndGroupByAgent(companyId string) map[string]int64
	CountDeploymentsByCompanyIdAndAgent(companyId, agentName string) int64
}
