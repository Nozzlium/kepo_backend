package helper

import (
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

func QuestionLikeParamToResponse(
	param param.QuestionLikeParam,
) response.QuestionLikeResponse {
	return response.QuestionLikeResponse{
		IsLiked:    param.IsLiked,
		QuestionID: param.QuestionLike.QuestionID,
		UserID:     param.QuestionLike.UserID,
	}
}
