package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type ingressService struct {
	ingressRepo repository.Ingress
}

func (i ingressService) GetById(id, agent, processId string) v1.Ingress {
	return i.ingressRepo.GetById(id, agent, processId)
}

func (i ingressService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Ingress, int64) {
	if ownerReference == "" {
		return i.ingressRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return i.ingressRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (i ingressService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Ingress {
	return i.ingressRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewIngressService returns service.Ingress type service
func NewIngressService(ingressRepo repository.Ingress) service.Ingress {
	return &ingressService{
		ingressRepo: ingressRepo,
	}
}
