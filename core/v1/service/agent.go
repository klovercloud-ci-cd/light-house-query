package service

import v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"

// Agent business operations.
type Agent interface {
	Get(companyId string) []v1.Agent
	GetK8sObjs(agent, processId string) v1.K8sObjsInfo
	GetPodsByCertificate(agent, processId, certificateId string) []v1.K8sPod
	GetPodsByClusterRole(agent, processId, clusterRoleId string) []v1.K8sPod
	GetPodsByClusterRoleBinding(agent, processId, clusterRoleBindingId string) []v1.K8sPod
	GetPodsByConfigMap(agent, processId, configMapId string) []v1.K8sPod
	GetPodsByDaemonSet(agent, processId, daemonSetId string) []v1.K8sPod
	GetPodsByDeployment(agent, processId, deploymentId string) []v1.K8sPod
	GetPodsByIngress(agent, processId, ingressId string) []v1.K8sPod
	GetPodsByNamespace(agent, processId, namespaceId string) []v1.K8sPod
	GetPodsByNetworkPolicy(agent, processId, networkPolicyId string) []v1.K8sPod
	GetPodsByNode(agent, processId, nodeId string) []v1.K8sPod
	GetPodsByPV(agent, processId, PVId string) []v1.K8sPod
	GetPodsByPVC(agent, processId, PVCId string) []v1.K8sPod
	GetPodsByReplicaSet(agent, processId, replicaSetId string) []v1.K8sPod
	GetPodsByRole(agent, processId, roleId string) []v1.K8sPod
	GetPodsByRoleBinding(agent, processId, roleBindingId string) []v1.K8sPod
	GetPodsBySecret(agent, processId, secretId string) []v1.K8sPod
	GetPodsByService(agent, processId, serviceId string) []v1.K8sPod
	GetPodsByServiceAccount(agent, processId, serviceAccountId string) []v1.K8sPod
	GetPodsByStatefulSet(agent, processId, statefulSetId string) []v1.K8sPod
}
