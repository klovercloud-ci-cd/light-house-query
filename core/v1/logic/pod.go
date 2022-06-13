package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type podService struct {
	podRepo        repository.Pod
	deploymentRepo repository.Deployment
}

func (p podService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Pod, int64) {
	if ownerReference == "" {
		return p.podRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return p.podRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (p podService) GetDashboardData(companyId string) v1.PodDashboardData {
	deploymentCountMap := p.deploymentRepo.CountDeploymentsByCompanyIdAndGroupByAgent(companyId)
	podCountMap := make(map[string]map[string]int64)
	for agentName, _ := range deploymentCountMap {
		count := p.podRepo.CountSucceededStatusByCompanyIdAndAgent(companyId, agentName)
		if statusMap, ok := podCountMap[agentName]; ok {
			if val, ok2 := statusMap[string(v1.PodSucceeded)]; ok2 {
				statusMap[string(v1.PodSucceeded)] = val + count
			} else {
				statusMap[string(v1.PodSucceeded)] = count
			}
			podCountMap[agentName] = statusMap
		} else {
			podCountMap[agentName] = map[string]int64{string(v1.PodSucceeded): count}
		}
	}
	for agentName, _ := range deploymentCountMap {
		nullContainerStatusesMap := make(map[string]int64)
		podShortDtos := p.podRepo.GetNullContainerStatusesMap(companyId, agentName)
		for _, each := range podShortDtos {
			status := string(each.Obj.Status.Phase)
			if val, ok := nullContainerStatusesMap[status]; ok {
				nullContainerStatusesMap[status] = val + 1
			} else {
				nullContainerStatusesMap[status] = 1
			}
		}
		for status, count := range nullContainerStatusesMap {
			if statusMap, ok := podCountMap[agentName]; ok {
				if val, ok2 := statusMap[status]; ok2 {
					statusMap[status] = val + count
				} else {
					statusMap[status] = count
				}
				podCountMap[agentName] = statusMap
			} else {
				podCountMap[agentName] = map[string]int64{status: count}
			}
		}
	}
	for agentName, _ := range deploymentCountMap {
		containerStatusesMap := make(map[string]int64)
		podShortDtos := p.podRepo.GetContainerStatusesMap(companyId, agentName)
		for _, each := range podShortDtos {
			running := true
			var status string
			for _, eachContainerStatus := range each.Obj.Status.ContainerStatuses {
				if eachContainerStatus.State.Terminated != nil {
					status = eachContainerStatus.State.Terminated.Reason
					running = false
					break
				} else if eachContainerStatus.State.Waiting != nil {
					status = eachContainerStatus.State.Waiting.Reason
					running = false
					if eachContainerStatus.State.Waiting.Reason == "ImagePullBackOff" || eachContainerStatus.State.Waiting.Reason == "CrashLoopBackOff" {
						break
					}
				}
			}
			if running {
				status = string(v1.PodRunning)
			}
			if val, ok := containerStatusesMap[status]; ok {
				containerStatusesMap[status] = val + 1
			} else {
				containerStatusesMap[status] = 1
			}
		}
		for status, count := range containerStatusesMap {
			if statusMap, ok := podCountMap[agentName]; ok {
				if val, ok2 := statusMap[status]; ok2 {
					statusMap[status] = val + count
				} else {
					statusMap[status] = count
				}
				podCountMap[agentName] = statusMap
			} else {
				podCountMap[agentName] = map[string]int64{status: count}
			}
		}
	}
	var dashboardData v1.PodDashboardData
	for agentName, count := range deploymentCountMap {
		podCountAgentDataDto := v1.PodCountAgentDataDto{
			Name: agentName,
			Deployment: struct {
				Count int64 `json:"count"`
			}{Count: count},
		}
		if statusMap, ok := podCountMap[agentName]; ok {
			podCountAgentDataDto.Pods = statusMap
		}
		dashboardData.Data.Agent = append(dashboardData.Data.Agent, podCountAgentDataDto)
	}
	return dashboardData
}

// NewPodService returns service.Pod type service
func NewPodService(podRepo repository.Pod, deploymentRepo repository.Deployment) service.Pod {
	return &podService{
		podRepo:        podRepo,
		deploymentRepo: deploymentRepo,
	}
}
