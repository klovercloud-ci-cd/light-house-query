package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
)

type agentApi struct {
	agentService service.Agent
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Agents by company id
// @Tags Agent
// @Produce json
// @Param companyId query string true "Company Id"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Agent{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents [GET]
func (a agentApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	data := a.agentService.Get(companyId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// Get... Get K8sObjs Api
// @Summary Get K8sObjs api
// @Description Api for getting all K8sObjs short info by agent name and process id
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=v1.k8sObjsInfo{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/k8sobjs [GET]
func (a agentApi) GetK8sObjs(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	data := a.agentService.GetK8sObjs(agent, processId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// Get... Get Pods by Certificate Api
// @Summary Get Pods by Certificate api
// @Description Api for getting all K8SPods by agent name, process id and Certificate id
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Certificate ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/certificates/{certificateId}/pods [GET]
func (a agentApi) GetPodsByCertificate(context echo.Context) error {
	agent := context.Param("agent")
	certificateId := context.Param("certificateId")
	processId := context.QueryParam("processId")
	data := a.agentService.GetPodsByCertificate(agent, processId, certificateId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByClusterRole(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	clusterRoleId := context.Param("clusterRoleId")
	data := a.agentService.GetPodsByClusterRole(agent, processId, clusterRoleId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByClusterRoleBinding(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	clusterRoleBindingId := context.Param("clusterRoleBindingId")
	data := a.agentService.GetPodsByClusterRoleBinding(agent, processId, clusterRoleBindingId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByConfigMap(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	configMapId := context.Param("configMapId")
	data := a.agentService.GetPodsByConfigMap(agent, processId, configMapId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByDaemonSet(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	daemonSetId := context.Param("daemonSetId")
	data := a.agentService.GetPodsByDaemonSet(agent, processId, daemonSetId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByDeployment(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	deploymentId := context.Param("deploymentId")
	data := a.agentService.GetPodsByDeployment(agent, processId, deploymentId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByIngress(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	ingressId := context.Param("ingressId")
	data := a.agentService.GetPodsByIngress(agent, processId, ingressId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByNamespace(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	namespaceId := context.Param("namespaceId")
	data := a.agentService.GetPodsByNamespace(agent, processId, namespaceId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByNetworkPolicy(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	networkPolicyId := context.Param("networkPolicyId")
	data := a.agentService.GetPodsByNetworkPolicy(agent, processId, networkPolicyId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByNode(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	nodeId := context.Param("nodeId")
	data := a.agentService.GetPodsByNode(agent, processId, nodeId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByPV(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	pvId := context.Param("pvId")
	data := a.agentService.GetPodsByPV(agent, processId, pvId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByPVC(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	pvcId := context.Param("pvcId")
	data := a.agentService.GetPodsByPVC(agent, processId, pvcId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByReplicaSet(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	replicaSetId := context.Param("replicaSetId")
	data := a.agentService.GetPodsByReplicaSet(agent, processId, replicaSetId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByRole(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	roleId := context.Param("roleId")
	data := a.agentService.GetPodsByRole(agent, processId, roleId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByRoleBinding(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	roleBindingId := context.Param("roleBindingId")
	data := a.agentService.GetPodsByRoleBinding(agent, processId, roleBindingId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsBySecret(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	secretId := context.Param("secretId")
	data := a.agentService.GetPodsBySecret(agent, processId, secretId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByService(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	serviceId := context.Param("serviceId")
	data := a.agentService.GetPodsByService(agent, processId, serviceId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByServiceAccount(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	serviceAccountId := context.Param("serviceAccountId")
	data := a.agentService.GetPodsByServiceAccount(agent, processId, serviceAccountId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

func (a agentApi) GetPodsByStatefulSet(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	statefulSetId := context.Param("statefulSetId")
	data := a.agentService.GetPodsByStatefulSet(agent, processId, statefulSetId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// NewAgentApi returns api.Agent type api
func NewAgentApi(agentService service.Agent) api.Agent {
	return &agentApi{
		agentService: agentService,
	}
}
