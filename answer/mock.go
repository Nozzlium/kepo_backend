package answer

import (
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/data/repository"
	"nozzlium/kepo_backend/data/requestbody"
	"nozzlium/kepo_backend/data/result"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var answerRepositoryMock = repository.AnswerRepositoryMock{Mock: &mock.Mock{}}
var answerService = AnswerServiceImpl{
	AnswerRepository: &answerRepositoryMock,
}

var answerController AnswerController = &AnswerControllerImpl{
	AnswerService: &answerService,
	Validator:     validator.New(),
}

var createAnswerBody = requestbody.Answer{
	QuestionID: 1,
	Content:    "test",
}

var createAnswerBodyExceedLength = requestbody.Answer{
	QuestionID: 1,
	Content: `
		Lorem ipsum dolor sit amet, consectetuer adipiscing elit. 
		Aenean commodo ligula eget dolor. Aenean massa. Cum sociis 
		natoque penatibus et magnis dis parturient montes, nascetur 
		ridiculus mus. Donec quam felis, ultricies nec, pellentesque 
		eu, pretium quis, sem. Nulla consequat massa quis enim. 
		Donec pede justo, fringilla vel, aliquet nec, vulputate eget, 
		arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, 
		justo. Nullam dictum felis eu pede mollis pretium. Integer 
		tincidunt. Cras dapibu
	`,
}

var expectedAnswers = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      10,
		IsLiked:    0,
		Content:    "test",
	},
	{
		ID:         2,
		QuestionID: 1,
		UserID:     2,
		Username:   "user2",
		Likes:      5,
		IsLiked:    1,
		Content:    "test",
	},
}

var expectedAnswersFromSameUser = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      10,
		IsLiked:    0,
	},
	{
		ID:         3,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      5,
		IsLiked:    1,
	},
}

var expectedAnswersFromSameQuestion = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      10,
		IsLiked:    0,
	},
	{
		ID:         3,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      5,
		IsLiked:    1,
	},
}

var expectedOneAnswer = []result.AnswerResult{
	{
		ID:         1,
		QuestionID: 1,
		UserID:     1,
		Username:   "user1",
		Likes:      10,
		IsLiked:    0,
	},
}

func mockReturnCreatedAnswer() *mock.Call {
	return answerRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		entity.Answer{
			ID:         1,
			UserID:     1,
			Content:    "test",
			QuestionID: 1,
		},
		nil,
	)
}

func mockCreateAnswerError() *mock.Call {
	return answerRepositoryMock.Mock.On("Insert",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(entity.Answer{}, gorm.ErrInvalidField)
}

func mockReturnAnswers() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedAnswers,
		nil,
	)
}

func mockReturnAnswersFromSameUser() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedAnswersFromSameUser,
		nil,
	)
}

func mockReturnOneAnswer() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedOneAnswer,
		nil,
	)
}

func mockReturnEmpty() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.AnswerResult{},
		nil,
	)
}

func mockReturnAnswersError() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		[]result.AnswerResult{},
		gorm.ErrInvalidDB,
	)
}

func mockReturnOneAnswerDetailed() *mock.Call {
	return answerRepositoryMock.Mock.On("FindOneDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedAnswers[0],
		nil,
	)
}

func mockReturnOneAnswerDetailedError() *mock.Call {
	return answerRepositoryMock.Mock.On("FindOneDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		result.AnswerResult{},
		gorm.ErrRecordNotFound,
	)
}

func mockReturnAnswerByQuestion() *mock.Call {
	return answerRepositoryMock.Mock.On("FindDetailed",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(
		expectedAnswersFromSameQuestion,
		nil,
	)
}
