package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	PodRouter(g.Group("/pods"))
	CertificateRouter(g.Group("/certificates"))
	NodeRouter(g.Group("/nodes"))
	ClusterRoleRouter(g.Group("/cluster_roles"))
}

// PodRouter api/v1/pods/* router
func PodRouter(g *echo.Group) {
	podApi := NewPodApi(dependency.GetV1PodService())
	g.GET("", podApi.Get)
}

// CertificateRouter api/v1/certificates/* router
func CertificateRouter(g *echo.Group) {
	certificateApi := NewCertificateApi(dependency.GetV1CertificateService())
	g.GET("", certificateApi.Get)
}

// NodeRouter api/v1/nodes/* router
func NodeRouter(g *echo.Group) {
	nodeApi := NewNodeApi(dependency.GetV1NodeService())
	g.GET("", nodeApi.Get)
}

// ClusterRoleRouter api/v1/cluster_roles/* router
func ClusterRoleRouter(g *echo.Group) {
	nodeApi := NewNodeApi(dependency.GetV1NodeService())
	g.GET("", nodeApi.Get)
}
