package api

import "github.com/labstack/echo/v4"

// ServiceAccount api operations
type ServiceAccount interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
}
