package api

import "github.com/labstack/echo/v4"

type Agent interface {
	Get(context echo.Context) error
}
