package v1

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetQueryOption(context echo.Context) v1.ResourceQueryOption {
	option := v1.ResourceQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	ascendingSort := context.QueryParam("sort")
	if page == "" {
		option.Pagination.Page = 0
		option.Pagination.Limit = 10
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}
	if ascendingSort == "true" {
		option.AscendingSort = true
	}
	return option
}
