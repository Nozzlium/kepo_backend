package repositoryimpl

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/result"

	"gorm.io/gorm"
)

type AnswerRepositoryImpl struct {
}

func (repository *AnswerRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error) {
	insert := DB.WithContext(ctx).Create(&answer)
	return answer, insert.Error
}

func (repository *AnswerRepositoryImpl) FindBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]entity.Answer, error) {
	answers := []entity.Answer{}
	find := DB.WithContext(ctx).
		Where(&param.Answer).
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Find(&answers)
	return answers, find.Error
}

func (repository *AnswerRepositoryImpl) FindOneBy(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (entity.Answer, error) {
	answer := entity.Answer{}
	find := DB.WithContext(ctx).
		Where(&param.Answer).
		First(&answer)
	return answer, find.Error
}

func (repository *AnswerRepositoryImpl) FindDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) ([]result.AnswerResult, error) {
	answers := []result.AnswerResult{}
	find := DB.WithContext(ctx).
		Table(
			`answers as a
			join users as u on u.id = a.user_id 
			left join answer_likes al on al.answer_id = a.id,
			left join answer_likes al1 on al1.answer_id = a.id and al1.user_id = ?`,
			param.UserID,
		).
		Select(
			`a.id,
			a.content,
			a.question_id,
			u.Uid as user_id,
			u.username as username,
			count(al.answer_id) as likes,
			al1.answer_id as user_liked`,
		)
	if param.Answer.QuestionID != 0 {
		find = find.Where("a.question_id = ?", param.Answer.QuestionID)
	}
	if param.Answer.UserID != 0 {
		find = find.Where("a.user_id = ?", param.Answer.UserID)
	}
	find = find.
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Find(&answers)
	return answers, find.Error
}

func (repository *AnswerRepositoryImpl) FindOneDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (result.AnswerResult, error) {
	answer := result.AnswerResult{}
	find := DB.WithContext(ctx).
		Table(
			`answers as a
			join users as u on u.id = a.user_id 
			left join answer_likes al on al.answer_id = a.id,
			left join answer_likes al1 on al1.answer_id = a.id and al1.user_id = ?`,
			param.UserID,
		).
		Select(
			`a.id,
			a.content,
			a.question_id,
			u.Uid as user_id,
			u.username as username,
			count(al.answer_id) as likes,
			al1.answer_id as user_liked`,
		)
	if param.Answer.QuestionID != 0 {
		find = find.Where("a.question_id = ?", param.Answer.QuestionID)
	}
	if param.Answer.UserID != 0 {
		find = find.Where("a.user_id = ?", param.Answer.UserID)
	}
	find = find.First(&answer)
	return answer, find.Error
}
