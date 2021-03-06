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
	AgentRouter(g.Group("/agents"))
}

// CertificateRouter api/v1/certificates/* router
func CertificateRouter(g *echo.Group) {
	certificateApi := NewCertificateApi(dependency.GetV1CertificateService())
	g.GET("", certificateApi.Get)
	g.GET("/:id", certificateApi.GetByID)
}

// ClusterRoleRouter api/v1/cluster-roles/* router
func ClusterRoleRouter(g *echo.Group) {
	clusterRoleApi := NewClusterRoleApi(dependency.GetV1ClusterRoleService())
	g.GET("", clusterRoleApi.Get)
	g.GET("/:id", clusterRoleApi.GetByID)

}

// ClusterRoleBindingRouter api/v1/cluster-role-bindings/* router
func ClusterRoleBindingRouter(g *echo.Group) {
	clusterRoleBindingApi := NewClusterRoleBindingApi(dependency.GetV1ClusterRoleBindingService())
	g.GET("", clusterRoleBindingApi.Get)
	g.GET("/:id", clusterRoleBindingApi.GetByID)
}

// ConfigMapRouter api/v1/config-maps/* router
func ConfigMapRouter(g *echo.Group) {
	configMapApi := NewConfigMapApi(dependency.GetV1ConfigMapService())
	g.GET("", configMapApi.Get)
	g.GET("/:id", configMapApi.GetByID)
}

// DaemonSetRouter api/v1/daemon-sets/* router
func DaemonSetRouter(g *echo.Group) {
	daemonSetApi := NewDaemonSetApi(dependency.GetV1DaemonSetService())
	g.GET("", daemonSetApi.Get)
	g.GET("/:id", daemonSetApi.GetByID)
}

// DeploymentRouter api/v1/deployments/* router
func DeploymentRouter(g *echo.Group) {
	deploymentApi := NewDeploymentApi(dependency.GetV1DeploymentService())
	g.GET("", deploymentApi.Get)
	g.GET("/:id", deploymentApi.GetByID)
}

// IngressRouter api/v1/ingresses/* router
func IngressRouter(g *echo.Group) {
	ingressApi := NewIngressApi(dependency.GetV1IngressService())
	g.GET("", ingressApi.Get)
	g.GET("/:id", ingressApi.GetByID)
}

// NamespaceRouter api/v1/namespaces/* router
func NamespaceRouter(g *echo.Group) {
	namespaceApi := NewNamespaceApi(dependency.GetV1NamespaceService())
	g.GET("", namespaceApi.Get)
	g.GET("/:id", namespaceApi.GetByID)
}

// NetworkPolicyRouter api/v1/network-policies/* router
func NetworkPolicyRouter(g *echo.Group) {
	networkPolicyApi := NewNetworkPolicyApi(dependency.GetV1NetworkPolicyService())
	g.GET("", networkPolicyApi.Get)
	g.GET("/:id", networkPolicyApi.GetByID)
}

// NodeRouter api/v1/nodes/* router
func NodeRouter(g *echo.Group) {
	nodeApi := NewNodeApi(dependency.GetV1NodeService())
	g.GET("", nodeApi.Get)
	g.GET("/:id", nodeApi.GetByID)
}

// PodRouter api/v1/pods/* router
func PodRouter(g *echo.Group) {
	podApi := NewPodApi(dependency.GetV1PodService())
	g.GET("", podApi.Get)
	g.GET("/:id", podApi.GetByID)
}

// PersistentVolumeRouter api/v1/persistent-volumes/* router
func PersistentVolumeRouter(g *echo.Group) {
	persistentVolumeApi := NewPersistentVolumeApi(dependency.GetV1PersistentVolumeService())
	g.GET("", persistentVolumeApi.Get)
	g.GET("/:id", persistentVolumeApi.GetByID)
}

// PersistentVolumeClaimRouter api/v1/persistent-volume-claims/* router
func PersistentVolumeClaimRouter(g *echo.Group) {
	persistentVolumeClaimApi := NewPersistentVolumeClaimApi(dependency.GetV1PersistentVolumeClaimService())
	g.GET("", persistentVolumeClaimApi.Get)
	g.GET("/:id", persistentVolumeClaimApi.GetByID)
}

// ReplicaSetRouter api/v1/replica-sets/* router
func ReplicaSetRouter(g *echo.Group) {
	replicaSetApi := NewReplicaSetApi(dependency.GetV1ReplicaSetService())
	g.GET("", replicaSetApi.Get)
	g.GET("/:id", replicaSetApi.GetByID)
}

// RoleRouter api/v1/roles/* router
func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService())
	g.GET("", roleApi.Get)
	g.GET("/:id", roleApi.GetByID)
}

// RoleBindingRouter api/v1/role-bindings/* router
func RoleBindingRouter(g *echo.Group) {
	roleBindingApi := NewRoleBindingApi(dependency.GetV1RoleBindingService())
	g.GET("", roleBindingApi.Get)
	g.GET("/:id", roleBindingApi.GetByID)
}

// SecretRouter api/v1/secrets/* router
func SecretRouter(g *echo.Group) {
	secretApi := NewSecretApi(dependency.GetV1SecretService())
	g.GET("", secretApi.Get)
	g.GET("/:id", secretApi.GetByID)
}

// ServiceRouter api/v1/services/* router
func ServiceRouter(g *echo.Group) {
	serviceApi := NewServiceApi(dependency.GetV1ServiceService())
	g.GET("", serviceApi.Get)
	g.GET("/:id", serviceApi.GetByID)
}

// ServiceAccountRouter api/v1/service-accounts/* router
func ServiceAccountRouter(g *echo.Group) {
	serviceAccountApi := NewServiceAccountApi(dependency.GetV1ServiceAccountService())
	g.GET("", serviceAccountApi.Get)
	g.GET("/:id", serviceAccountApi.GetByID)
}

// StatefulSetRouter api/v1/stateful-sets/* router
func StatefulSetRouter(g *echo.Group) {
	statefulSetApi := NewStatefulSetApi(dependency.GetV1StatefulSetService())
	g.GET("", statefulSetApi.Get)
	g.GET("/:id", statefulSetApi.GetByID)
}

// AgentRouter api/v1/agents/* router
func AgentRouter(g *echo.Group) {
	agentApi := NewAgentApi(dependency.GetV1AgentService())
	g.GET("", agentApi.Get)
	g.GET("/:agent", agentApi.GetByID)
	g.GET("/:agent/k8sobjs", agentApi.GetK8sObjs)
	g.GET("/:agent/daemonSets/:daemonSetId/pods", agentApi.GetPodsByDaemonSet)
	g.GET("/:agent/deployments/:deploymentId/pods", agentApi.GetPodsByDeployment)
	g.GET("/:agent/replicaSets/:replicaSetId/pods", agentApi.GetPodsByReplicaSet)
	g.GET("/:agent/statefulSets/:statefulSetId/pods", agentApi.GetPodsByStatefulSet)
}
