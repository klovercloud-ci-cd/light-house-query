package api

import "github.com/labstack/echo/v4"

// ClusterRoleBinding api operations
type ClusterRoleBinding interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
	GetByProcessId(context echo.Context) error
}
