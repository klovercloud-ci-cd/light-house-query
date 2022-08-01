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

// Get... Get by ID Api
// @Summary Get by ID api
// @Description Api for getting an Agent by name and company id
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param companyId query string true "Company Id"
// @Success 200 {object} common.ResponseDTO{data=v1.Agent{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent} [GET]
func (a agentApi) GetByID(context echo.Context) error {
	agent := context.Param("agent")
	companyId := context.QueryParam("companyId")
	data := a.agentService.GetByName(agent, companyId)
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

// Get... Get Pods by DaemonSet Api
// @Summary Get Pods by DaemonSet api
// @Description Api for getting all K8SPods by agent name, process id and DaemonSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Certificate ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/daemonSets/{daemonSetId}/pods [GET]
func (a agentApi) GetPodsByDaemonSet(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	daemonSetId := context.Param("daemonSetId")
	data := a.agentService.GetPodsByDaemonSet(agent, processId, daemonSetId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// Get... Get Pods by Deployment Api
// @Summary Get Pods by Deployment api
// @Description Api for getting all K8SPods by agent name, process id and Deployment uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Deployment ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/deployments/{deploymentId}/pods [GET]
func (a agentApi) GetPodsByDeployment(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	deploymentId := context.Param("deploymentId")
	data := a.agentService.GetPodsByDeployment(agent, processId, deploymentId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// Get... Get Pods by ReplicaSet Api
// @Summary Get Pods by ReplicaSet api
// @Description Api for getting all K8SPods by agent name, process id and ReplicaSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "ReplicaSet ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/replicaSets/{replicaSetId}/pods [GET]
func (a agentApi) GetPodsByReplicaSet(context echo.Context) error {
	agent := context.Param("agent")
	processId := context.QueryParam("processId")
	replicaSetId := context.Param("replicaSetId")
	data := a.agentService.GetPodsByReplicaSet(agent, processId, replicaSetId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// Get... Get Pods by StatefulSet Api
// @Summary Get Pods by StatefulSet api
// @Description Api for getting all K8SPods by agent name, process id and StatefulSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "StatefulSet ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/statefulSets/{statefulSetId}/pods [GET]
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
