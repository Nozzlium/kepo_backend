package helper

import (
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/data/result"
)

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
		Likes:     result.Likes,
		IsLiked:   result.IsLiked != 0,
		CreatedAt: TimeToString(result.CreatedAt),
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
