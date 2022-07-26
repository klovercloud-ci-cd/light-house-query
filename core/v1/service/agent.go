package service

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
)

// Agent business operations.
type Agent interface {
	Get(companyId string) []v1.Agent
	GetByName(agent, companyId string) v1.Agent
	GetK8sObjs(agent, processId string) v1.K8sObjsInfo
	GetPodsByDaemonSet(agent, processId, daemonSetId string) []v1.K8sPod
	GetPodsByDeployment(agent, processId, deploymentId string) []v1.K8sPod
	GetPodsByReplicaSet(agent, processId, replicaSetId string) []v1.K8sPod
	GetPodsByStatefulSet(agent, processId, statefulSetId string) []v1.K8sPod
}
