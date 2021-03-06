package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleBindingService struct {
	roleBindingRepo repository.RoleBinding
}

func (r roleBindingService) GetById(id, agent, processId string) v1.RoleBinding {
	return r.roleBindingRepo.GetById(id, agent, processId)
}

func (r roleBindingService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64) {
	if ownerReference == "" {
		return r.roleBindingRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return r.roleBindingRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (r roleBindingService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.RoleBinding {
	return r.roleBindingRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewRoleBindingService returns service.RoleBinding type service
func NewRoleBindingService(roleBindingRepo repository.RoleBinding) service.RoleBinding {
	return &roleBindingService{
		roleBindingRepo: roleBindingRepo,
	}
}
