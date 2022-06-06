package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type persistentVolumeService struct {
	persistentVolumeRepo repository.PersistentVolume
}

func (p persistentVolumeService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolume, int64) {
	if ownerReference == "" {
		return p.persistentVolumeRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return p.persistentVolumeRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

// NewPersistentVolumeService returns service.PersistentVolume type service
func NewPersistentVolumeService(persistentVolumeRepo repository.PersistentVolume) service.PersistentVolume {
	return &persistentVolumeService{
		persistentVolumeRepo: persistentVolumeRepo,
	}
}
