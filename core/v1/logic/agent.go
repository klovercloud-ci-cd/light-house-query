package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type agentService struct {
	agentRepo                    repository.Agent
	certificateService           service.Certificate
	clusterRoleService           service.ClusterRole
	clusterRoleBindingService    service.ClusterRoleBinding
	configMapService             service.ConfigMap
	daemonSetService             service.DaemonSet
	deploymentService            service.Deployment
	ingressService               service.Ingress
	namespaceService             service.Namespace
	networkPolicyService         service.NetworkPolicy
	nodeService                  service.Node
	podService                   service.Pod
	persistentVolumeService      service.PersistentVolume
	persistentVolumeClaimService service.PersistentVolumeClaim
	replicaSetService            service.ReplicaSet
	roleService                  service.Role
	roleBindingService           service.RoleBinding
	secretService                service.Secret
	serviceService               service.Service
	serviceAccountService        service.ServiceAccount
	statefulSetService           service.StatefulSet
}

func (a agentService) Get(companyId string) []v1.Agent {
	return a.agentRepo.Get(companyId)
}

func (a agentService) GetPodsByDaemonSet(agent, processId, daemonSetId string) []v1.K8sPod {
	daemonSet := a.daemonSetService.GetById(daemonSetId, agent, processId)
	pods := a.podService.GetByAgentAndProcessIdAndLabels(agent, processId, daemonSet.Obj.Spec.Selector.MatchLabels)
	var k8sPods []v1.K8sPod
	for _, each := range pods {
		k8sPods = append(k8sPods, each.Obj)
	}
	return k8sPods
}

func (a agentService) GetPodsByDeployment(agent, processId, deploymentId string) []v1.K8sPod {
	deployment := a.deploymentService.GetById(deploymentId, agent, processId)
	pods := a.podService.GetByAgentAndProcessIdAndLabels(agent, processId, deployment.Obj.Spec.Selector.MatchLabels)
	var k8sPods []v1.K8sPod
	for _, each := range pods {
		k8sPods = append(k8sPods, each.Obj)
	}
	return k8sPods
}

func (a agentService) GetPodsByReplicaSet(agent, processId, replicaSetId string) []v1.K8sPod {
	replicaSet := a.replicaSetService.GetById(replicaSetId, agent, processId)
	pods := a.podService.GetByAgentAndProcessIdAndLabels(agent, processId, replicaSet.Obj.Spec.Selector.MatchLabels)
	var k8sPods []v1.K8sPod
	for _, each := range pods {
		k8sPods = append(k8sPods, each.Obj)
	}
	return k8sPods
}

func (a agentService) GetPodsByStatefulSet(agent, processId, statefulSetId string) []v1.K8sPod {
	statefulSet := a.statefulSetService.GetById(statefulSetId, agent, processId)
	pods := a.podService.GetByAgentAndProcessIdAndLabels(agent, processId, statefulSet.Obj.Spec.Selector.MatchLabels)
	var k8sPods []v1.K8sPod
	for _, each := range pods {
		k8sPods = append(k8sPods, each.Obj)
	}
	return k8sPods
}

func (a agentService) GetK8sObjs(agent, processId string) v1.K8sObjsInfo {
	data := v1.K8sObjsInfo{}
	certificates := a.certificateService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var certificatesInfo []v1.K8sObjShortInfo
	for _, each := range certificates {
		certificatesInfo = append(certificatesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Certificates = certificatesInfo
	clusterRoles := a.clusterRoleService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var clusterRolesInfo []v1.K8sObjShortInfo
	for _, each := range clusterRoles {
		clusterRolesInfo = append(clusterRolesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.ClusterRoles = clusterRolesInfo
	clusterRoleBindings := a.clusterRoleBindingService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var clusterRoleBindingsInfo []v1.K8sObjShortInfo
	for _, each := range clusterRoleBindings {
		clusterRoleBindingsInfo = append(clusterRoleBindingsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.ClusterRoleBindings = clusterRoleBindingsInfo
	configMaps := a.configMapService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var configMapsInfo []v1.K8sObjShortInfo
	for _, each := range configMaps {
		configMapsInfo = append(configMapsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.ConfigMaps = configMapsInfo
	daemonSets := a.daemonSetService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var daemonSetInfo []v1.DaemonSetShortInfo
	for _, each := range daemonSets {
		daemonSetInfo = append(daemonSetInfo, v1.DaemonSetShortInfo{
			Name:              each.Obj.Name,
			Namespace:         each.Obj.Namespace,
			UID:               string(each.Obj.UID),
			NumberReady:       each.Obj.Status.NumberReady,
			NumberAvailable:   each.Obj.Status.NumberAvailable,
			NumberUnavailable: each.Obj.Status.NumberUnavailable,
		})
	}
	data.DaemonSets = daemonSetInfo
	deployments := a.deploymentService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var deploymentsInfo []v1.DeploymentShortInfo
	for _, each := range deployments {
		deploymentsInfo = append(deploymentsInfo, v1.DeploymentShortInfo{
			Name:                each.Obj.Name,
			Namespace:           each.Obj.Namespace,
			UID:                 string(each.Obj.UID),
			Replicas:            each.Obj.Status.Replicas,
			AvailableReplicas:   each.Obj.Status.AvailableReplicas,
			UnavailableReplicas: each.Obj.Status.UnavailableReplicas,
			ReadyReplicas:       each.Obj.Status.ReadyReplicas,
		})
	}
	data.Deployments = deploymentsInfo
	ingresses := a.ingressService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var ingressesInfo []v1.K8sObjShortInfo
	for _, each := range ingresses {
		ingressesInfo = append(ingressesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Ingresses = ingressesInfo
	namespaces := a.namespaceService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var namespacesInfo []v1.K8sObjShortInfo
	for _, each := range namespaces {
		namespacesInfo = append(namespacesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Namespaces = namespacesInfo
	networkPolicies := a.networkPolicyService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var networkPoliciesInfo []v1.K8sObjShortInfo
	for _, each := range networkPolicies {
		networkPoliciesInfo = append(networkPoliciesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.NetworkPolicies = networkPoliciesInfo
	nodes := a.nodeService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var nodesInfo []v1.K8sObjShortInfo
	for _, each := range nodes {
		nodesInfo = append(nodesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Nodes = nodesInfo
	pvs := a.persistentVolumeService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var pvsInfo []v1.K8sObjShortInfo
	for _, each := range pvs {
		pvsInfo = append(pvsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.PersistentVolumes = pvsInfo
	pvcs := a.persistentVolumeClaimService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var pvcsInfo []v1.K8sObjShortInfo
	for _, each := range pvcs {
		pvcsInfo = append(pvcsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.PersistentVolumeClaims = pvcsInfo
	replicaSets := a.replicaSetService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var replicaSetsInfo []v1.ReplicaSetShortInfo
	for _, each := range replicaSets {
		replicaSetsInfo = append(replicaSetsInfo, v1.ReplicaSetShortInfo{
			Name:              each.Obj.Name,
			Namespace:         each.Obj.Namespace,
			UID:               string(each.Obj.UID),
			Replicas:          each.Obj.Status.Replicas,
			AvailableReplicas: each.Obj.Status.AvailableReplicas,
			ReadyReplicas:     each.Obj.Status.ReadyReplicas,
		})
	}
	data.ReplicaSets = replicaSetsInfo
	roles := a.roleService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var rolesInfo []v1.K8sObjShortInfo
	for _, each := range roles {
		rolesInfo = append(rolesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Roles = rolesInfo
	roleBindings := a.roleBindingService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var roleBindingsInfo []v1.K8sObjShortInfo
	for _, each := range roleBindings {
		roleBindingsInfo = append(roleBindingsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.RoleBindings = roleBindingsInfo
	secrets := a.secretService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var secretsInfo []v1.K8sObjShortInfo
	for _, each := range secrets {
		secretsInfo = append(secretsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Secrets = secretsInfo
	services := a.serviceService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var servicesInfo []v1.K8sObjShortInfo
	for _, each := range services {
		servicesInfo = append(servicesInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.Services = servicesInfo
	serviceAccounts := a.serviceAccountService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var serviceAccountsInfo []v1.K8sObjShortInfo
	for _, each := range serviceAccounts {
		serviceAccountsInfo = append(serviceAccountsInfo, v1.K8sObjShortInfo{
			Name:      each.Obj.Name,
			Namespace: each.Obj.Namespace,
			UID:       string(each.Obj.UID),
		})
	}
	data.ServiceAccounts = serviceAccountsInfo
	statefulSets := a.statefulSetService.GetByAgentAndProcessIdWithoutPagination(agent, processId)
	var statefulSetsInfo []v1.StatefulSetShortInfo
	for _, each := range statefulSets {
		statefulSetsInfo = append(statefulSetsInfo, v1.StatefulSetShortInfo{
			Name:          each.Obj.Name,
			Namespace:     each.Obj.Namespace,
			UID:           string(each.Obj.UID),
			Replicas:      each.Obj.Status.Replicas,
			ReadyReplicas: each.Obj.Status.ReadyReplicas,
		})
	}
	data.StatefulSets = statefulSetsInfo
	return data
}

// NewAgentService returns service.Agent type service
func NewAgentService(agentRepo repository.Agent, certificateService service.Certificate, clusterRoleService service.ClusterRole,
	clusterRoleBindingService service.ClusterRoleBinding, configMapService service.ConfigMap, daemonSetService service.DaemonSet,
	deploymentService service.Deployment, ingressService service.Ingress, namespaceService service.Namespace,
	networkPolicyService service.NetworkPolicy, nodeService service.Node, podService service.Pod, persistentVolumeService service.PersistentVolume,
	persistentVolumeClaimService service.PersistentVolumeClaim, replicaSetService service.ReplicaSet, roleService service.Role,
	roleBindingService service.RoleBinding, secretService service.Secret, serviceService service.Service, serviceAccountService service.ServiceAccount,
	statefulSetService service.StatefulSet) service.Agent {
	return &agentService{
		agentRepo:                    agentRepo,
		certificateService:           certificateService,
		clusterRoleService:           clusterRoleService,
		clusterRoleBindingService:    clusterRoleBindingService,
		configMapService:             configMapService,
		daemonSetService:             daemonSetService,
		deploymentService:            deploymentService,
		ingressService:               ingressService,
		namespaceService:             namespaceService,
		networkPolicyService:         networkPolicyService,
		nodeService:                  nodeService,
		podService:                   podService,
		persistentVolumeService:      persistentVolumeService,
		persistentVolumeClaimService: persistentVolumeClaimService,
		replicaSetService:            replicaSetService,
		roleService:                  roleService,
		roleBindingService:           roleBindingService,
		secretService:                secretService,
		serviceService:               serviceService,
		serviceAccountService:        serviceAccountService,
		statefulSetService:           statefulSetService,
	}
}
