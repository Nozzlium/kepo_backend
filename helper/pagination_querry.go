package helper

import (
	"net/http"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/data/param"
	"strconv"
)

func GetPaginationParamFromQuerry(
	request *http.Request,
) param.PaginationParam {
	pagination := param.InitPaginationParam()

	queries := request.URL.Query()
	pageNo, err := strconv.Atoi(queries.Get(constants.PAGE_NO))
	if err == nil && pageNo > 0 {
		pagination.PageNo = pageNo
	}

	pageSize, err := strconv.Atoi(queries.Get(constants.PAGE_SIZE))
	if err == nil {
		pagination.PageSize = pageSize
	}

	keyword := queries.Get(constants.KEYWORD)
	pagination.Keyword = keyword

	sortBy := queries.Get(constants.SORT_BY)
	if sortBy != "" {
		pagination.SortBy = sortBy
	}

	order := queries.Get(constants.ORDER)
	if order != "" {
		pagination.Order = order
	}

	return pagination
}
