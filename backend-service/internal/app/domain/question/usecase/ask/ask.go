package ask

import (
	"backend-service/internal/app/domain/question/entity"
	"context"
)

type AskService struct {
}

func NewAskService() *AskService {
	return &AskService{}
}

func (i *AskService) Execute(ctx context.Context, req Request) (Response, error) {
	_, err := req.MapIntoQuestion()
	if err != nil {
		return Response{}, err
	}

	return Response{
		Answer: "TODO: answer question",
	}, nil
}

type Request struct {
	Question string `json:"question" validate:"required"`
}

type Response struct {
	Answer string `json:"answer"`
}

func (i *Request) MapIntoQuestion() (entity.Question, error) {
	var question entity.Question = entity.Question{
		Question: i.Question,
	}
	if err := question.Validate(); err != nil {
		return entity.Question{}, err
	}
	return question, nil
}
