package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/param"
	"nozzlium/kepo_backend/data/result"

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
	find := DB.WithContext(ctx).Model(&entity.Question{}).
		Table(`questions 
			join users u on u.id = questions.user_id
			join categories c on c.id = questions.category_id
			left join question_likes ql on questions.id = ql.question_id
			left join question_likes qll on qll.question_id = questions.id and qll.user_id = ?
			left join answers a on a.question_id  = questions.id`,
			param.UserID,
		).
		Select(
			`questions.id,
			questions.content,
			questions.description,
			u.id as user_id,
			u.username as username,
			c.id as category_id,
			c.name as category_name,
			count(distinct ql.user_id) as likes,
			COUNT(distinct a.id) as answers,
			qll.question_id as user_liked,
			questions.created_at,
			questions.updated_at
			`,
		)
	if param.Keyword != "" {
		find = find.Where(
			`MATCH(questions.content, questions.description)
				AGAINST(? IN NATURAL LANGUAGE MODE)
			`,
			param.Keyword,
		)
	}
	if param.Question.UserID != 0 {
		find = find.Where(
			"questions.user_id = ?", param.Question.UserID,
		)
	}
	if param.Question.ID != 0 {
		find = find.Where(
			"questions.id = ?", param.Question.ID,
		)
	}
	if param.Question.CategoryID != 0 {
		find = find.Where(
			"questions.category_id", param.Question.CategoryID,
		)
	}
	find = find.Group("questions.id").
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Order("questions.created_at DESC").Find(&res)
	return res, find.Error
}

func (repository *QuestionRepositoryImpl) FindOneDetailedBy(ctx context.Context, DB *gorm.DB, param param.QuestionParam) (result.QuestionResult, error) {
	question := result.QuestionResult{}
	find := DB.WithContext(ctx).Model(&entity.Question{}).
		Table(`questions 
			join users u on u.id = questions.user_id
			join categories c on c.id = questions.category_id
			left join question_likes ql on questions.id = ql.question_id
			left join question_likes qll on qll.question_id = questions.id and qll.user_id = ?
			left join answers a on a.question_id  = questions.id`,
			param.UserID,
		).
		Select(
			`questions.id,
			questions.content,
			questions.description,
			u.id as user_id,
			u.username as username,
			c.id as category_id,
			c.name as category_name,
			count(distinct ql.user_id) as likes,
			COUNT(distinct a.id) as answers,
			qll.question_id as user_liked,
			questions.created_at,
			questions.updated_at
			`,
		)
	if param.Question.ID != 0 {
		find = find.Where(
			"questions.id = ?", param.Question.ID,
		)
	}
	if param.Question.UserID != 0 {
		find = find.Where(
			"questions.user_id = ?", param.Question.UserID,
		)
	}
	find = find.Group("questions.id").First(&question)
	return question, find.Error
}

func (repository *QuestionRepositoryImpl) FindDetailedLikedByUser(ctx context.Context, DB *gorm.DB, param param.LikedQuestionParam) ([]result.QuestionResult, error) {
	questions := []result.QuestionResult{}
	find := DB.WithContext(ctx).
		Table(`question_likes ql 
				join questions q on ql.question_id = q.id 
				join users u ON u.id = q.user_id 
				join categories c ON c.id = q.category_id 
				left join question_likes ql2 on ql2.question_id = ql.question_id
				left JOIN question_likes ql3 on ql3.question_id = ql.question_id  and ql3.user_id = ?
				left join answers a on ql.question_id = a.question_id`,
			param.UserID).
		Select(`ql.question_id as id,
				u.id as user_id ,
				u.username ,
				c.id as category_id,
				c.name as category_name,
				q.content ,
				q.description,
				count(DISTINCT ql2.user_id) as likes,
				count(DISTINCT a.id) as answers,
				ql3.question_id as user_liked,
				q.created_at,
				q.updated_at
			`)
	find = find.Group("ql.question_id, ql.created_at").Where("ql.user_id = ? AND q.deleted_at IS NULL", param.LikerID).
		Limit(param.PageSize).
		Offset((param.PageNo - 1) * param.PageSize).
		Order("ql.created_at DESC").
		Find(&questions)
	return questions, find.Error
}

func (repository *QuestionRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, question entity.Question) (entity.Question, error) {
	delete := DB.WithContext(ctx).Delete(&question)
	return question, delete.Error
}

func (repository *QuestionRepositoryImpl) Update(ctx context.Context, DB *gorm.DB, question entity.Question) (entity.Question, error) {
	save := DB.WithContext(ctx).Model(&question).Updates(question)
	return question, save.Error
}
