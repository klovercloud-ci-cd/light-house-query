package api

import "github.com/labstack/echo/v4"

// Service api operations
type Service interface {
	Get(context echo.Context) error
}
