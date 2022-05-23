package v1

import (
	"github.com/klovercloud-ci-cd/light-house-query/api/common"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/api"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type persistentVolumeApi struct {
	persistentVolumeService service.PersistentVolume
}

// Get... Get Api
// @Summary Get api
// @Description Api for getiing all Persistent Volume Claims by agent name
// @Tags PersistentVolume
// @Produce json
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO{data=[]v1.PersistentVolume{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/persistentVolumes [GET]
func (p persistentVolumeApi) Get(context echo.Context) error {
	agent := context.QueryParam("agent")
	option := GetQueryOption(context)
	data, total := p.persistentVolumeService.Get(agent, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data,
		&metadata, "Successful")
}

// NewPersistentVolumeApi returns api.PersistentVolume type api
func NewPersistentVolumeApi(persistentVolumeService service.PersistentVolume) api.PersistentVolume {
	return &persistentVolumeApi{
		persistentVolumeService: persistentVolumeService,
	}
}
