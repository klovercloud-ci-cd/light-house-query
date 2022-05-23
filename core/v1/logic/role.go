package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleService struct {
	roleServiceRepo repository.Role
}

func (r roleService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Role, int64) {
	return r.roleServiceRepo.Get(agent, option)
}

// NewRoleService returns service.RoleService type service
func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{
		roleServiceRepo: roleRepo,
	}
}
