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

func (p podService) GetById(id, agent, processId string) v1.Pod {
	return p.podRepo.GetById(id, agent, processId)
}

func (p podService) GetByAgentAndProcessIdAndLabels(agent, processId string, labels map[string]string) []v1.Pod {
	return p.podRepo.GetByAgentAndProcessIdAndLabels(agent, processId, labels)
}

func (p podService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Pod, int64) {
	if ownerReference == "" {
		return p.podRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return p.podRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (p podService) GetDashboardData(companyId, agentName string) v1.PodDashboardData {
	deploymentCount := p.deploymentRepo.CountDeploymentsByCompanyIdAndAgent(companyId, agentName)

	podStatusMap := make(map[string]int64)

	podStatusMap[string(v1.PodSucceeded)] = p.podRepo.CountSucceededStatusByCompanyIdAndAgent(companyId, agentName)

	podStatusMap = p.GetNullContainerStatusesMap(companyId, agentName, podStatusMap)

	podStatusMap = p.GetContainerStatusesMap(companyId, agentName, podStatusMap)

	return v1.PodDashboardData{
		Agent: v1.PodCountAgentDataDto{
			Name: agentName,
			Pods: podStatusMap,
			Deployment: struct {
				Count int64 `json:"count"`
			}(struct {
				Count int64
			}{
				Count: deploymentCount,
			})},
	}
}

func (p podService) GetNullContainerStatusesMap(companyId, agentName string, podStatusMap map[string]int64) map[string]int64 {
	podShortDtos := p.podRepo.GetNullContainerStatusesMap(companyId, agentName)
	for _, each := range podShortDtos {
		status := string(each.Obj.Status.Phase)
		if val, ok := podStatusMap[status]; ok {
			podStatusMap[status] = val + 1
		} else {
			podStatusMap[status] = 1
		}
	}
	return podStatusMap
}

func (p podService) GetContainerStatusesMap(companyId, agentName string, podStatusMap map[string]int64) map[string]int64 {
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
		if val, ok := podStatusMap[status]; ok {
			podStatusMap[status] = val + 1
		} else {
			podStatusMap[status] = 1
		}
	}
	return podStatusMap
}

// NewPodService returns service.Pod type service
func NewPodService(podRepo repository.Pod, deploymentRepo repository.Deployment) service.Pod {
	return &podService{
		podRepo:        podRepo,
		deploymentRepo: deploymentRepo,
	}
}
