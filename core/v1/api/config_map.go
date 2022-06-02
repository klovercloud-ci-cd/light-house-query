package api

import "github.com/labstack/echo/v4"

// ConfigMap api operations
type ConfigMap interface {
	Get(context echo.Context) error
}
