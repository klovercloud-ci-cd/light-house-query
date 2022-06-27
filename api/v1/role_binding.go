package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type roleBindingApi struct {
	roleBindingService service.RoleBinding
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Role Bindings by agent name, owner reference and process id
// @Tags RoleBinding
// @Produce json
// @Param owner-reference path string false "Owner Reference"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO{data=[]v1.RoleBinding{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/role-bindings [GET]
func (r roleBindingApi) Get(context echo.Context) error {
	agent := context.QueryParam("agent")
	ownerReference := context.QueryParam("owner-reference")
	processId := context.QueryParam("processId")
	option := GetQueryOption(context)
	data, total := r.roleBindingService.Get(agent, ownerReference, processId, option)
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

// NewRoleBindingApi returns api.RoleBinding type api
func NewRoleBindingApi(roleBindingService service.RoleBinding) api.RoleBinding {
	return &roleBindingApi{
		roleBindingService: roleBindingService,
	}
}
