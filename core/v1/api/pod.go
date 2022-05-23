package api

import "github.com/labstack/echo/v4"

// Pod api operations
type Pod interface {
	Get(context echo.Context) error
}
