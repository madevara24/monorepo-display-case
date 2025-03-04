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

	// Create vector for the question
	vector, err := i.openAIClient.CreateEmbedding(ctx, req.Question)
	if err != nil {
		return Response{}, err
	}

	// Find similar content
	similar, err := i.questionRepo.FindSimilar(ctx, vector, req.Limit)
	if err != nil {
		return Response{}, err
	}

	if len(similar) == 0 {
		return Response{}, fmt.Errorf("no similar content found")
	}

	// Extract contents for context
	var contents []string
	for _, s := range similar {
		contents = append(contents, s.Content)
	}

	// Generate response using similar content
	answer, err := i.openAIClient.GenerateResponse(ctx, req.Question, contents)
	if err != nil {
		return Response{}, err
	}

	res := Response{
		Answer: answer,
	}
	res.MapIntoSet(similar)

	return res, nil
}

type Request struct {
	Question string `json:"question" validate:"required"`
	Limit    int    `json:"limit"`
}

type Response struct {
	Answer string `json:"answer"`
	Set    []Set  `json:"set"`
}

type Set struct {
	Content string  `json:"content"`
	Score   float64 `json:"score"`
}

func (i *Response) MapIntoSet(similar []entity.Embedding) {
	set := []Set{}
	for _, s := range similar {
		set = append(set, Set{Content: s.Content, Score: s.Score})
	}

	i.Set = set
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
