package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/result"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnswerRepositoryImpl struct {
}

func NewAnswerRepository() *AnswerRepositoryImpl {
	return &AnswerRepositoryImpl{}
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
		Model(&entity.Answer{}).
		Table(
			`answers
			join users as u on u.id = answers.user_id
			join questions as q on q.id = answers.question_id
			left join answer_likes al on al.answer_id = answers.id
			left join answer_likes al1 on al1.answer_id = answers.id and al1.user_id = ?`,
			param.UserID,
		).
		Select(
			`answers.id,
			answers.content,
			answers.question_id,
			q.user_id as question_poster_id,
			u.id as user_id,
			u.username as username,
			count(distinct al.answer_id) as likes,
			al1.answer_id as user_liked,
			answers.created_at
			`,
		)
	if param.Answer.QuestionID != 0 {
		find = find.Where("answers.question_id = ?", param.Answer.QuestionID)
	}

	order := param.Order

	var sortBy string
	switch param.SortBy {
	case "DTE":
		sortBy = "answers.created_at"
	case "LKE":
		sortBy = "likes"
	default:
		sortBy = "answers.created_at"
	}

	if param.Answer.UserID != 0 {
		find = find.Where("answers.user_id = ?", param.Answer.UserID)
	}
	find = find.Group("answers.id").
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Order(
			clause.OrderByColumn{
				Column: clause.Column{
					Name: sortBy,
				},
				Desc: order == "DESC",
			},
		).
		Find(&answers)
	return answers, find.Error
}

func (repository *AnswerRepositoryImpl) FindOneDetailed(ctx context.Context, DB *gorm.DB, param param.AnswerParam) (result.AnswerResult, error) {
	answer := result.AnswerResult{}
	find := DB.WithContext(ctx).Model(&entity.Answer{}).
		Table(
			`answers
			join users as u on u.id = answers.user_id 
			join questions as q on q.id = answers.question_id
			left join answer_likes al on al.answer_id = answers.id
			left join answer_likes al1 on al1.answer_id = answers.id and al1.user_id = ?`,
			param.UserID,
		).
		Select(
			`answers.id,
			answers.content,
			answers.question_id,
			q.user_id as question_poster_id,
			u.id as user_id,
			u.username as username,
			count(distinct al.answer_id) as likes,
			al1.answer_id as user_liked,
			answers.created_at
			`,
		)
	if param.Answer.ID != 0 {
		find = find.Where(
			"answers.id = ?", param.Answer.ID,
		)
	}
	if param.Answer.QuestionID != 0 {
		find = find.Where("answers.question_id = ?", param.Answer.QuestionID)
	}
	if param.Answer.UserID != 0 {
		find = find.Where("answers.user_id = ?", param.Answer.UserID)
	}
	find = find.Group("answers.id").First(&answer)
	return answer, find.Error
}

func (repository *AnswerRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error) {
	delete := DB.WithContext(ctx).Delete(&answer)
	return answer, delete.Error
}

func (repository *AnswerRepositoryImpl) Update(ctx context.Context, DB *gorm.DB, answer entity.Answer) (entity.Answer, error) {
	save := DB.WithContext(ctx).Model(&answer).Updates(answer)
	return answer, save.Error
}
