package api

import "github.com/labstack/echo/v4"

// PersistentVolume api operations
type PersistentVolume interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
}
