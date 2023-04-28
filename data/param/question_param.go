package param

import "nozzlium/kepo_backend/data/entity"

type QuestionParam struct {
	PaginationParam
	UserID   uint
	Question entity.Question
}

func InitQuestionParam() QuestionParam {
	return QuestionParam{
		PaginationParam: InitPaginationParam(),
	}
}
