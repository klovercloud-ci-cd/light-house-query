package api

import (
	"github.com/labstack/echo/v4"
)

// DaemonSet api operations
type DaemonSet interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
	GetByProcessId(context echo.Context) error
}
