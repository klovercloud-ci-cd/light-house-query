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
	PersistentVolumeRouter(g.Group("/persistent-volumeRouter"))
	ClusterRoleRouter(g.Group("/cluster-roles"))
	ClusterRoleRouter(g.Group("/cluster-role-bindings"))
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

// ClusterRoleRouter api/v1/cluster-roles/* router
func ClusterRoleRouter(g *echo.Group) {
	clusterRoleApi := NewClusterRoleApi(dependency.GetV1ClusterRoleService())
	g.GET("", clusterRoleApi.Get)
}

// ClusterRoleBindingRouter api/v1/cluster-role-bindings/* router
func ClusterRoleBindingRouter(g *echo.Group) {
	clusterRoleBindingApi := NewClusterRoleBindingApi(dependency.GetV1ClusterRoleBindingService())
	g.GET("", clusterRoleBindingApi.Get)
}

// ConfigMapRouter api/v1/config-maps/* router
func ConfigMapRouter(g *echo.Group) {
	configMapApi := NewConfigMapApi(dependency.GetV1ConfigMapService())
	g.GET("", configMapApi.Get)
}

// PersistentVolumeRouter api/v1/persistent-volume/* router
func PersistentVolumeRouter(g *echo.Group) {
	persistentVolumeApi := NewPersistentVolumeApi(dependency.GetV1PersistentVolumeService())
	g.GET("", persistentVolumeApi.Get)
}
