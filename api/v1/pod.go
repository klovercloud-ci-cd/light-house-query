package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type podApi struct {
	podService service.Pod
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Pods by agent name, owner reference and process id
// @Tags Pod
// @Produce json
// @Param owner-reference path string false "Owner Reference"
// @Param action string true "action [dashboard_data]"
// @Param companyId string false "Company Id when action is dashboard_data"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Pod{}, v1.DashboardData{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/pods [GET]
func (p podApi) Get(context echo.Context) error {
	action := context.QueryParam("action")
	agent := context.QueryParam("agent")
	if action == "dashboard_data" {
		companyId := context.QueryParam("companyId")
		data := p.podService.GetDashboardData(companyId, agent)
		return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
	}
	ownerReference := context.QueryParam("owner-reference")
	processId := context.QueryParam("processId")
	option := GetQueryOption(context)
	data, total := p.podService.Get(agent, ownerReference, processId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?agent=" + context.QueryParam("agent") + "&action=" + context.QueryParam("action") + "&companyId" + context.QueryParam("companyId") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?agent=" + context.QueryParam("agent") + "&action=" + context.QueryParam("action") + "&companyId" + context.QueryParam("companyId") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?agent=" + context.QueryParam("agent") + "&action=" + context.QueryParam("action") + "&companyId" + context.QueryParam("companyId") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data,
		&metadata, "Successful")
}

// Get... Get by ID Api
// @Summary Get by ID api
// @Description Api for getting a Pod by id, agent name, and process id
// @Tags Pod
// @Produce json
// @Param id query string true "ID"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Success 200 {object} common.ResponseDTO{data=v1.Pod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/pods/{id} [GET]
func (p podApi) GetByID(context echo.Context) error {
	id := context.Param("id")
	agent := context.QueryParam("agent")
	processId := context.QueryParam("processId")
	data := p.podService.GetById(id, agent, processId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// NewPodApi returns api.Pod type api
func NewPodApi(podService service.Pod) api.Pod {
	return &podApi{
		podService: podService,
	}
}
