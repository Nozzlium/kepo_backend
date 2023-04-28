package param

import "nozzlium/kepo_backend/data/entity"

type QuestionLikeParam struct {
	QuestionLike entity.QuestionLike
	IsLiked      bool
}
