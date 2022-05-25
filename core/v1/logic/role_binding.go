package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleBindingService struct {
	roleBindingRepo repository.RoleBinding
}

func (r roleBindingService) Get(agent string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64) {
	return r.roleBindingRepo.Get(agent, option)
}

func (r roleBindingService) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64) {
	return r.roleBindingRepo.GetByOwnerReference(agent, ownerReference, option)
}

// NewRoleBindingService returns service.RoleBinding type service
func NewRoleBindingService(roleBindingRepo repository.RoleBinding) service.RoleBinding {
	return &roleBindingService{
		roleBindingRepo: roleBindingRepo,
	}
}
