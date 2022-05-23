package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	PodRouter(g.Group("/pods"))
}

// PodRouter api/v1/pods/* router
func PodRouter(g *echo.Group) {
	podApi := NewPodApi(dependency.GetV1PodService())
	g.GET("", podApi.Get)
}
