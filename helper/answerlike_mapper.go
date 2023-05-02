package helper

import (
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/response"
)

func AnswerLikeParamToResponse(
	param param.AnswerLikeParam,
) response.AnswerLikeResponse {
	return response.AnswerLikeResponse{
		IsLike:   param.IsLike,
		AnswerID: param.AnswerLike.AnswerID,
		UserID:   param.AnswerLike.UserID,
	}
}
