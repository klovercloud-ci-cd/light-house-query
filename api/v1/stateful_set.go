package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type statefulSetApi struct {
	statefulSetService service.StatefulSet
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all Stateful sets by agent name, owner reference and process id
// @Tags StatufulSet
// @Produce json
// @Param owner-reference path string false "Owner Reference"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO{data=[]v1.StatefulSet{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/stateful-sets [GET]
func (s statefulSetApi) Get(context echo.Context) error {
	agent := context.QueryParam("agent")
	ownerReference := context.QueryParam("owner-reference")
	processId := context.QueryParam("processId")
	option := GetQueryOption(context)
	data, total := s.statefulSetService.Get(agent, ownerReference, processId, option)
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
// @Description Api for getting a Statuful Set by id, agent name, and process id
// @Tags StatufulSet
// @Produce json
// @Param id query string true "ID"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Success 200 {object} common.ResponseDTO{data=v1.StatufulSet{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/stateful-sets/{id} [GET]
func (s statefulSetApi) GetByID(context echo.Context) error {
	id := context.Param("id")
	agent := context.QueryParam("agent")
	processId := context.QueryParam("processId")
	data := s.statefulSetService.GetById(id, agent, processId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
}

// NewStatefulSetApi returns api.StatefulSet type api
func NewStatefulSetApi(statefulSetService service.StatefulSet) api.StatefulSet {
	return &statefulSetApi{
		statefulSetService: statefulSetService,
	}
}
