package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Routes base router
func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud light house query service")
}
