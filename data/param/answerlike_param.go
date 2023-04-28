package param

import "nozzlium/kepo_backend/data/entity"

type AnswerLikeParam struct {
	AnswerLike entity.AnswerLike
	IsLike     bool
}
