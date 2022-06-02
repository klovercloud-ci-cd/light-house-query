package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type persistentVolumeClaimService struct {
	persistentVolumeClaimRepo repository.PersistentVolumeClaim
}

func (p persistentVolumeClaimService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64) {
	return p.persistentVolumeClaimRepo.Get(agent, ownerReference, processId, option)
}

// NewPersistentVolumeClaimService returns service.PersistentVolumeClaim type service
func NewPersistentVolumeClaimService(persistentVolumeClaimRepo repository.PersistentVolumeClaim) service.PersistentVolumeClaim {
	return &persistentVolumeClaimService{
		persistentVolumeClaimRepo: persistentVolumeClaimRepo,
	}
}
