package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type persistentVolumeClaimService struct {
	persistentVolumeClaimRepo repository.PersistentVolumeClaim
}

func (p persistentVolumeClaimService) GetById(id, agent, processId string) v1.PersistentVolumeClaim {
	return p.persistentVolumeClaimRepo.GetById(id, agent, processId)
}

func (p persistentVolumeClaimService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64) {
	if ownerReference == "" {
		return p.persistentVolumeClaimRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return p.persistentVolumeClaimRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (p persistentVolumeClaimService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.PersistentVolumeClaim {
	return p.persistentVolumeClaimRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewPersistentVolumeClaimService returns service.PersistentVolumeClaim type service
func NewPersistentVolumeClaimService(persistentVolumeClaimRepo repository.PersistentVolumeClaim) service.PersistentVolumeClaim {
	return &persistentVolumeClaimService{
		persistentVolumeClaimRepo: persistentVolumeClaimRepo,
	}
}
