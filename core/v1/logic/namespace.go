package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type namespaceService struct {
	namespaceRepo repository.Namespace
}

func (n namespaceService) GetById(id, agent, processId string) v1.Namespace {
	return n.namespaceRepo.GetById(id, agent, processId)
}

func (n namespaceService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Namespace, int64) {
	if ownerReference == "" {
		return n.namespaceRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return n.namespaceRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (n namespaceService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Namespace {
	return n.namespaceRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewNamespaceService returns service.Namespace type service
func NewNamespaceService(namespaceRepo repository.Namespace) service.Namespace {
	return &namespaceService{
		namespaceRepo: namespaceRepo,
	}
}
