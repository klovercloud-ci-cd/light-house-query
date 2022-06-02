package api

import "github.com/labstack/echo/v4"

// PersistentVolumeClaim api operations
type PersistentVolumeClaim interface {
	Get(context echo.Context) error
}
