package ask

import (
	"backend-service/internal/app/domain/question/entity"
	"backend-service/internal/app/domain/question/repository"
	"backend-service/internal/pkg/integration/openai"
	"context"
	"fmt"
)

type AskUsecase struct {
	openAIClient *openai.Client
	questionRepo *repository.QuestionRepository
}

func NewAskUsecase(openAIClient *openai.Client, questionRepo *repository.QuestionRepository) *AskUsecase {
	return &AskUsecase{
		openAIClient: openAIClient,
		questionRepo: questionRepo,
	}
}

func (i *AskUsecase) Execute(ctx context.Context, req Request) (Response, error) {
	_, err := req.MapIntoQuestion()
	if err != nil {
		return Response{}, err
	}

	// Create embedding for the question
	embedding, err := i.openAIClient.CreateEmbedding(ctx, req.Question)
	if err != nil {
		return Response{}, err
	}

	// Find similar content
	similar, err := i.questionRepo.FindSimilar(ctx, embedding, 5)
	if err != nil {
		return Response{}, err
	}

	fmt.Println(similar)

	// TODO: Use similar content to generate response with OpenAI

	return Response{
		// Answer: "TODO: answer question",
		Answer: similar[0].Content,
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
		Content: i.Question,
	}
	if err := question.Validate(); err != nil {
		return entity.Question{}, err
	}
	return question, nil
}
