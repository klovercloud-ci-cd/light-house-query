package api

import "github.com/labstack/echo/v4"

// RoleBinding api operations
type RoleBinding interface {
	Get(context echo.Context) error
}
