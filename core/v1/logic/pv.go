package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type persistentVolumeService struct {
	persistentVolumeRepo repository.PersistentVolume
}

func (p persistentVolumeService) GetById(id, agent, processId string) v1.PersistentVolume {
	return p.persistentVolumeRepo.GetById(id, agent, processId)
}

func (p persistentVolumeService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolume, int64) {
	if ownerReference == "" {
		return p.persistentVolumeRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return p.persistentVolumeRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (p persistentVolumeService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.PersistentVolume {
	return p.persistentVolumeRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewPersistentVolumeService returns service.PersistentVolume type service
func NewPersistentVolumeService(persistentVolumeRepo repository.PersistentVolume) service.PersistentVolume {
	return &persistentVolumeService{
		persistentVolumeRepo: persistentVolumeRepo,
	}
}
