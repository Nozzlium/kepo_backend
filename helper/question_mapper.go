package helper

import (
	"nozzlium/kepo_backend/data/response"
	"nozzlium/kepo_backend/data/result"
)

func QuestionResultToResponse(
	result result.QuestionResult,
) response.QuestionResponse {
	return response.QuestionResponse{
		ID: result.ID,
		User: response.UserResponse{
			ID:       result.UserID,
			Username: result.Username,
		},
		Category: response.CategoryResponse{
			ID:   result.CategoryID,
			Name: result.CategoryName,
		},
		Content:     result.Content,
		Description: result.Description,
		Likes:       result.Likes,
		Answers:     result.Answers,
		IsLiked:     result.UserLiked != 0,
	}
}

func QuestionResultsToResponses(
	results []result.QuestionResult,
) []response.QuestionResponse {
	responses := []response.QuestionResponse{}
	for _, result := range results {
		responses = append(
			responses, QuestionResultToResponse(result),
		)
	}
	return responses
}
