package api

import "github.com/labstack/echo/v4"

// Deployment api operations
type Deployment interface {
	Get(context echo.Context) error
}
