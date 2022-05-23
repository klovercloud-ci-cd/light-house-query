package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type podService struct {
	podRepo repository.Pod
}

func (p podService) Get(agent string) ([]v1.Pod, int64) {
	return p.podRepo.Get(agent)
}

// NewPodService returns service.Pod type service
func NewPodService(podRepo service.Pod) service.Pod {
	return &podService{
		podRepo: podRepo,
	}
}
