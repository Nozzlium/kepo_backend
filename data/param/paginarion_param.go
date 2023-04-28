package param

type PaginationParam struct {
	PageNo   int
	PageSize int
	Keyword  string
}

func InitPaginationParam() PaginationParam {
	return PaginationParam{
		PageNo:   1,
		PageSize: 10,
	}
}
