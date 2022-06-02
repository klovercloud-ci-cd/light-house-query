package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleBindingService struct {
	roleBindingRepo repository.RoleBinding
}

func (r roleBindingService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64) {
	return r.roleBindingRepo.Get(agent, ownerReference, processId, option)
}

// NewRoleBindingService returns service.RoleBinding type service
func NewRoleBindingService(roleBindingRepo repository.RoleBinding) service.RoleBinding {
	return &roleBindingService{
		roleBindingRepo: roleBindingRepo,
	}
}
