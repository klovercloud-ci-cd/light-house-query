package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Pod business operations.
type Pod interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Pod, int64)
	GetDashboardData(companyId, agentName string) v1.PodDashboardData
	GetByAgentAndProcessIdAndLabels(agent, processId string, labels map[string]string) []v1.Pod
}
