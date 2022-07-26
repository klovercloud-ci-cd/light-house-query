package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type networkPolicyApi struct {
	networkPolicyService service.NetworkPolicy
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Network Policies by agent name, owner reference and process
// @Tags Certificate
// @Produce json
// @Param owner-reference path string false "Owner Reference"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO{data=[]v1.NetworkPolicy{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/network-policies [GET]
func (n networkPolicyApi) Get(context echo.Context) error {
	agent := context.QueryParam("agent")
	ownerReference := context.QueryParam("owner-reference")
	processId := context.QueryParam("processId")
	option := GetQueryOption(context)
	data, total := n.networkPolicyService.Get(agent, ownerReference, processId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?agent=" + context.QueryParam("agent") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?agent=" + context.QueryParam("agent") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?agent=" + context.QueryParam("agent") + "&owner-reference=" + context.QueryParam("owner-reference") + "&processId=" + context.QueryParam("processId") + "&sort=" + context.QueryParam("sort") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data,
		&metadata, "Successful")
}

// Get... Get by ID Api
// @Summary Get by ID api
// @Description Api for getting a NetworkPolicy by id, agent name, and process id
// @Tags NetworkPolicy
// @Produce json
// @Param id query string true "ID"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Success 200 {object} common.ResponseDTO{data=v1.NetworkPolicy{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/network-policies/{id} [GET]
func (n networkPolicyApi) GetByID(context echo.Context) error {
	id := context.Param("id")
	agent := context.QueryParam("agent")
	processId := context.QueryParam("processId")
	data := n.networkPolicyService.GetById(id, agent, processId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// NewNetworkPolicyApi returns api.NetworkPolicy type api
func NewNetworkPolicyApi(networkPolicyService service.NetworkPolicy) api.NetworkPolicy {
	return &networkPolicyApi{
		networkPolicyService: networkPolicyService,
	}
}
