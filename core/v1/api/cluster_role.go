package api

import "github.com/labstack/echo/v4"

// ClusterRole api operations
type ClusterRole interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
	GetByProcessId(context echo.Context) error
}
