package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type ingressService struct {
	ingressRepo repository.Ingress
}

func (i ingressService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Ingress, int64) {
	return i.ingressRepo.Get(agent, option)
}

func (i ingressService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Ingress, int64) {
	return i.ingressRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewIngressService returns service.Ingress type service
func NewIngressService(ingressRepo repository.Ingress) service.Ingress {
	return &ingressService{
		ingressRepo: ingressRepo,
	}
}
