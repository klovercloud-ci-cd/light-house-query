package api

import "github.com/labstack/echo/v4"

// Namespace api operations
type Namespace interface {
	Get(context echo.Context) error
}
