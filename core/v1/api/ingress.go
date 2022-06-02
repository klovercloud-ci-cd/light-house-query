package api

import "github.com/labstack/echo/v4"

// Ingress api operations
type Ingress interface {
	Get(context echo.Context) error
}
