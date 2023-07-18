package helper

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/data/result"
)

func AnswerEntityToResponse(
	entity entity.Answer,
) response.AnswerResponse {
	return response.AnswerResponse{
		ID:         entity.ID,
		QuestionID: entity.QuestionID,
		User: response.UserResponse{
			ID: entity.UserID,
		},
		Content: entity.Content,
		Likes:   0,
		IsLiked: false,
	}
}

func AnswerResultToResponse(
	result result.AnswerResult,
) response.AnswerResponse {
	return response.AnswerResponse{
		ID:         result.ID,
		Content:    result.Content,
		QuestionID: result.QuestionID,
		User: response.UserResponse{
			ID:       result.UserID,
			Username: result.Username,
		},
		Likes:   result.Likes,
		IsLiked: result.IsLiked != 0,
	}
}

func AnswersResultSliceToResponsesSlice(
	results []result.AnswerResult,
) []response.AnswerResponse {
	res := []response.AnswerResponse{}
	for _, result := range results {
		res = append(res, AnswerResultToResponse(result))
	}
	return res
}
