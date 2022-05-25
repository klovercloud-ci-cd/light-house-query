package api

import "github.com/labstack/echo/v4"

// StatefulSet api operations
type StatefulSet interface {
	Get(context echo.Context) error
	GetByOwnerReference(context echo.Context) error
}
