package param

type PaginationParam struct {
	PageNo   int
	PageSize int
	Keyword  string
	SortBy   string
	Order    string
}

func InitPaginationParam() PaginationParam {
	return PaginationParam{
		PageNo:   1,
		PageSize: 10,
		SortBy:   "DTE",
		Order:    "DESC",
	}
}
