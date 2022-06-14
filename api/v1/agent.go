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

// NewAgentApi returns api.Agent type api
func NewAgentApi(agentService service.Agent) api.Agent {
	return &agentApi{
		agentService: agentService,
	}
}
