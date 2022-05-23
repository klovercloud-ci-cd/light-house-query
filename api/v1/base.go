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
	PersistentVolumeRouter(g.Group("/persistent-volumes"))
	PersistentVolumeCLaimRouter(g.Group("/persistent-volume-claims"))
	ReplicaSetRouter(g.Group("/replica-sets"))
	RoleRouter(g.Group("/roles"))
	ClusterRoleRouter(g.Group("/cluster-roles"))
	ClusterRoleBindingRouter(g.Group("/cluster-role-bindings"))
	ConfigMapRouter(g.Group("/config-maps"))
	DaemonSetRouter(g.Group("/daemon-sets"))
	DeploymentRouter(g.Group("/deployments"))
	IngressRouter(g.Group("/ingresses"))
	NamespaceRouter(g.Group("/namespaces"))
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

// DaemonSetRouter api/v1/daemon-sets/* router
func DaemonSetRouter(g *echo.Group) {
	daemonSetApi := NewDaemonSetApi(dependency.GetV1DaemonSetService())
	g.GET("", daemonSetApi.Get)
}

// PersistentVolumeRouter api/v1/persistent-volume/* router
func PersistentVolumeRouter(g *echo.Group) {
	persistentVolumeApi := NewPersistentVolumeApi(dependency.GetV1PersistentVolumeService())
	g.GET("", persistentVolumeApi.Get)
}

// PersistentVolumeCLaimRouter api/v1/persistent-volume-claims/* router
func PersistentVolumeCLaimRouter(g *echo.Group) {
	persistentVolumeClaimApi := NewPersistentVolumeClaimApi(dependency.GetV1PersistentVolumeClaimService())
	g.GET("", persistentVolumeClaimApi.Get)
}

// DeploymentRouter api/v1/deployments/* router
func DeploymentRouter(g *echo.Group) {
	deploymentApi := NewDeploymentApi(dependency.GetV1DeploymentService())
	g.GET("", deploymentApi.Get)
}

// NamespaceRouter api/v1/namespaces/* router
func NamespaceRouter(g *echo.Group) {
	namespaceApi := NewNamespaceApi(dependency.GetV1NamespaceService())
	g.GET("", namespaceApi.Get)
}

// IngressRouter api/v1/ingresses/* router
func IngressRouter(g *echo.Group) {
	ingressApi := NewIngressApi(dependency.GetV1IngressService())
	g.GET("", ingressApi.Get)
}

// ReplicaSetRouter api/v1/replica-sets/* router
func ReplicaSetRouter(g *echo.Group) {
	replicaSetApi := NewReplicaSetApi(dependency.GetV1ReplicaSetService())
	g.GET("", replicaSetApi.Get)
}

// RoleRouter api/v1/roles/* router
func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService())
	g.GET("", roleApi.Get)
}
