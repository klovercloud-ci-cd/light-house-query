package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleService struct {
	roleRepo repository.Role
}

func (r roleService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Role, int64) {
	return r.roleRepo.Get(agent, option)
}

func (r roleService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Role, int64) {
	return r.roleRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewRoleService returns service.Role type service
func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{
		roleRepo: roleRepo,
	}
}
