package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	CertificateRouter(g.Group("/certificates"))
	ClusterRoleRouter(g.Group("/cluster-roles"))
	ClusterRoleBindingRouter(g.Group("/cluster-role-bindings"))
	ConfigMapRouter(g.Group("/config-maps"))
	DaemonSetRouter(g.Group("/daemon-sets"))
	DeploymentRouter(g.Group("/deployments"))
	IngressRouter(g.Group("/ingresses"))
	NamespaceRouter(g.Group("/namespaces"))
	NetworkPolicyRouter(g.Group("/network-policies"))
	NodeRouter(g.Group("/nodes"))
	PodRouter(g.Group("/pods"))
	PersistentVolumeRouter(g.Group("/persistent-volumes"))
	PersistentVolumeClaimRouter(g.Group("/persistent-volume-claims"))
	ReplicaSetRouter(g.Group("/replica-sets"))
	RoleRouter(g.Group("/roles"))
	RoleBindingRouter(g.Group("/role-bindings"))
	SecretRouter(g.Group("/secrets"))
	ServiceRouter(g.Group("/services"))
	ServiceAccountRouter(g.Group("/service-accounts"))
	StatefulSetRouter(g.Group("/stateful-sets"))
}

// CertificateRouter api/v1/certificates/* router
func CertificateRouter(g *echo.Group) {
	certificateApi := NewCertificateApi(dependency.GetV1CertificateService())
	g.GET("", certificateApi.Get)
	g.GET("/:owner-reference", certificateApi.GetByOwnerReference)
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

// PersistentVolumeRouter api/v1/persistent-volumes/* router
func PersistentVolumeRouter(g *echo.Group) {
	persistentVolumeApi := NewPersistentVolumeApi(dependency.GetV1PersistentVolumeService())
	g.GET("", persistentVolumeApi.Get)
}

// PersistentVolumeClaimRouter api/v1/persistent-volume-claims/* router
func PersistentVolumeClaimRouter(g *echo.Group) {
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

// RoleBindingRouter api/v1/role-bindings/* router
func RoleBindingRouter(g *echo.Group) {
	roleBindingApi := NewRoleBindingApi(dependency.GetV1RoleBindingService())
	g.GET("", roleBindingApi.Get)
}

// RoleRouter api/v1/roles/* router
func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService())
	g.GET("", roleApi.Get)
}

// NetworkPolicyRouter api/v1/network-policies/* router
func NetworkPolicyRouter(g *echo.Group) {
	networkPolicyApi := NewNetworkPolicyApi(dependency.GetV1NetworkPolicyService())
	g.GET("", networkPolicyApi.Get)
}

// NodeRouter api/v1/nodes/* router
func NodeRouter(g *echo.Group) {
	nodeApi := NewNodeApi(dependency.GetV1NodeService())
	g.GET("", nodeApi.Get)
}

// PodRouter api/v1/pods/* router
func PodRouter(g *echo.Group) {
	podApi := NewPodApi(dependency.GetV1PodService())
	g.GET("", podApi.Get)
}

// StatefulSetRouter api/v1/stateful-sets/* router
func StatefulSetRouter(g *echo.Group) {
	statefulSetApi := NewStatefulSetApi(dependency.GetV1StatefulSetService())
	g.GET("", statefulSetApi.Get)
}

// ServiceAccountRouter api/v1/service-accounts/* router
func ServiceAccountRouter(g *echo.Group) {
	serviceAccountApi := NewServiceAccountApi(dependency.GetV1ServiceAccountService())
	g.GET("", serviceAccountApi.Get)
}

// ServiceRouter api/v1/services/* router
func ServiceRouter(g *echo.Group) {
	serviceApi := NewServiceApi(dependency.GetV1ServiceService())
	g.GET("", serviceApi.Get)
}

// SecretRouter api/v1/secrets/* router
func SecretRouter(g *echo.Group) {
	secretApi := NewSecretApi(dependency.GetV1SecretService())
	g.GET("", secretApi.Get)
}
