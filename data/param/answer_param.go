package param

import "nozzlium/kepo_backend/data/entity"

type AnswerParam struct {
	PaginationParam
	UserID uint
	Answer entity.Answer
}
