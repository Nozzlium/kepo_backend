package repositoryimpl

import (
	"context"
	"fmt"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/repository/result"

	"gorm.io/gorm"
)

type QuestionRepositoryImpl struct {
}

func NewQuestionRepository() *QuestionRepositoryImpl {
	return &QuestionRepositoryImpl{}
}

func (repository *QuestionRepositoryImpl) Insert(ctx context.Context, DB *gorm.DB, question entity.Question) (entity.Question, error) {
	insert := DB.WithContext(ctx).
		Create(&question)
	return question, insert.Error
}

func (repository *QuestionRepositoryImpl) Find(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]entity.Question, error) {
	questions := []entity.Question{}
	find := DB.WithContext(ctx).
		Where(&param.Question).
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Find(&questions)

	return questions, find.Error
}

func (repository *QuestionRepositoryImpl) FindOneBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (entity.Question, error) {
	question := entity.Question{}
	find := DB.WithContext(ctx).
		Where(param.Question).
		First(&question)
	return question, find.Error
}

func (repository *QuestionRepositoryImpl) FindDetailed(ctx context.Context, DB *gorm.DB, param param.QuestionParam) ([]result.QuestionResult, error) {
	res := []result.QuestionResult{}
	fmt.Println(param)
	find := DB.WithContext(ctx).
		Table(`questions q 
			join users u on u.id = q.user_id
			join categories c on c.id = q.category_id
			left join question_likes ql on q.id = ql.question_id
			left join question_likes qll on qll.question_id = q.id and qll.user_id = ?
			left join answers a on a.question_id  = q.id`,
			param.UserID,
		).
		Select(
			`q.id,
			q.content,
			q.description,
			u.id as user_id,
			u.username as username,
			c.id as category_id,
			c.name as category_name,
			count(ql.question_id) as likes,
			COUNT(a.id) as answers,
			qll.question_id as user_liked
			`,
		)
	if param.Keyword != "" {
		find = find.Where(
			`MATCH(q.content, q.description)
				AGAINST(? IN NATURAL LANGUAGE MODE)
			`,
			param.Keyword,
		)
	}
	if param.Question.UserID != 0 {
		find = find.Where(
			"q.user_id = ?", param.Question.UserID,
		)
	}
	if param.Question.ID != 0 {
		find = find.Where(
			"q.id = ?", param.Question.ID,
		)
	}
	if param.Question.CategoryID != 0 {
		find = find.Where(
			"q.category_id", param.Question.CategoryID,
		)
	}
	find = find.Group("q.id").
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).Find(&res)
	return res, find.Error
}

func (repository *QuestionRepositoryImpl) FindOneDetailedBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (result.QuestionResult, error) {
	question := result.QuestionResult{}
	find := DB.WithContext(ctx).
		Table(`questions q 
			join users u on u.id = q.user_id
			join categories c on c.id = q.category_id
			left join question_likes ql on q.id = ql.question_id
			left join question_likes qll on qll.question_id = q.id and qll.user_id = ?
			left join answers a on a.question_id  = q.id`,
			param.UserID,
		).
		Select(
			`q.id,
			q.content,
			q.description,
			u.id as user_id,
			u.username as username,
			c.id as category_id,
			c.name as category_name,
			count(ql.question_id) as likes,
			COUNT(a.id) as answers,
			qll.question_id as user_liked
			`,
		)
	if param.Question.ID != 0 {
		find = find.Where(
			"q.id = ?", param.Question.ID,
		)
	}
	if param.Question.UserID != 0 {
		find = find.Where(
			"q.user_id = ?", param.Question.UserID,
		)
	}
	find = find.First(&question)
	return question, find.Error
}
