package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type roleBindingService struct {
	roleBindingServiceRepo repository.RoleBinding
}

func (r roleBindingService) Get(agent string, option v1.ResourceQueryOption) ([]v1.RoleBinding, int64) {
	return r.roleBindingServiceRepo.Get(agent, option)
}

// NewRoleBindingService returns service.RoleBindingService type service
func NewRoleBindingService(roleBindingRepo repository.RoleBinding) service.RoleBinding {
	return &roleBindingService{
		roleBindingServiceRepo: roleBindingRepo,
	}
}