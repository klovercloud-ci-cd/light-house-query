package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type deploymentService struct {
	deploymentRepo repository.Deployment
}

func (d deploymentService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	return d.deploymentRepo.Get(agent, option)
}

func (d deploymentService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	return d.deploymentRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewDeploymentService returns service.Deployment type service
func NewDeploymentService(deploymentRepo repository.Deployment) service.Deployment {
	return &deploymentService{
		deploymentRepo: deploymentRepo,
	}
}
