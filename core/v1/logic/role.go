package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleService struct {
	roleRepo repository.Role
}

func (r roleService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Role, int64) {
	return r.roleRepo.Get(agent, ownerReference, processId, option)
}

// NewRoleService returns service.Role type service
func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{
		roleRepo: roleRepo,
	}
}
