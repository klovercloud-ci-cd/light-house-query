package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleService struct {
	roleRepo repository.Role
}

func (r roleService) GetById(id, agent, processId string) v1.Role {
	return r.roleRepo.GetById(id, agent, processId)
}

func (r roleService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64) {
	if ownerReference == "" {
		return r.roleRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return r.roleRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (r roleService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Role {
	return r.roleRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewRoleService returns service.Role type service
func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{
		roleRepo: roleRepo,
	}
}
