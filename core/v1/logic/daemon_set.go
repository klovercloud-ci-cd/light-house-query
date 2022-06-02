package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type daemonSetService struct {
	daemonSetRepo repository.DaemonSet
}

func (d daemonSetService) Get(agent string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64) {
	return d.daemonSetRepo.Get(agent, option)
}

func (d daemonSetService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64) {
	return d.daemonSetRepo.GetByOwnerReference(agent, ownerReference, option)
}

func (d daemonSetService) GetByProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.DaemonSet, int64) {
	return d.daemonSetRepo.GetByOwnerReference(agent, processId, option)
}

// NewDaemonSetService returns service.Certificate type service
func NewDaemonSetService(daemonSetRepo repository.DaemonSet) service.DaemonSet {
	return &daemonSetService{
		daemonSetRepo: daemonSetRepo,
	}
}
