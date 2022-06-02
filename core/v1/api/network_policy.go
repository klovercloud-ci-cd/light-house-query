package api

import "github.com/labstack/echo/v4"

// NetworkPolicy api operations
type NetworkPolicy interface {
	Get(context echo.Context) error
}
