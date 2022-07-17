package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type deploymentService struct {
	deploymentRepo repository.Deployment
}

func (d deploymentService) GetById(id, agent, processId string) v1.Deployment {
	return d.deploymentRepo.GetById(id, agent, processId)
}

func (d deploymentService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	if ownerReference == "" {
		return d.deploymentRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return d.deploymentRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (d deploymentService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Deployment {
	return d.deploymentRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewDeploymentService returns service.Deployment type service
func NewDeploymentService(deploymentRepo repository.Deployment) service.Deployment {
	return &deploymentService{
		deploymentRepo: deploymentRepo,
	}
}
