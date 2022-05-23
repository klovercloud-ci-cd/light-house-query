package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
)

type clusterRoleApi struct {
	clusterRoleService service.ClusterRole
}

func (c clusterRoleApi) Get(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// NewClusterRoleApi returns api.ClusterRole type api
func NewClusterRoleApi(clusterRoleService service.ClusterRole) api.ClusterRole {
	return &clusterRoleApi{
		clusterRoleService: clusterRoleService,
	}
}
