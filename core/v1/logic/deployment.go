package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type deploymentService struct {
	deploymentRepo repository.Deployment
}

func (d deploymentService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	return d.deploymentRepo.Get(agent, ownerReference, processId, option)
}

// NewDeploymentService returns service.Deployment type service
func NewDeploymentService(deploymentRepo repository.Deployment) service.Deployment {
	return &deploymentService{
		deploymentRepo: deploymentRepo,
	}
}
