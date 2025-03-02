package store

import (
	"backend-service/internal/app/domain/question/entity"
	"backend-service/internal/app/domain/question/repository"
	"backend-service/internal/pkg/integration/openai"
	"context"
	"time"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	openAIClient *openai.Client
	questionRepo *repository.QuestionRepository
}

func NewStoreUsecase(openAIClient *openai.Client, questionRepo *repository.QuestionRepository) *StoreUsecase {
	return &StoreUsecase{
		openAIClient: openAIClient,
		questionRepo: questionRepo,
	}
}

func (i *StoreUsecase) Execute(ctx context.Context, req Request) error {
	embeddings, err := req.MapIntoEmbedding()
	if err != nil {
		return err
	}

	// Create embedding for the question
	vector, err := i.openAIClient.CreateEmbedding(ctx, req.Content)
	if err != nil {
		return err
	}

	embeddings.Embedding = vector

	err = i.questionRepo.StoreEmbedding(ctx, embeddings)
	if err != nil {
		return err
	}

	return nil
}

type Request struct {
	Category    string `json:"category" validate:"required"`
	Granularity string `json:"granularity" validate:"required"`
	Content     string `json:"content" validate:"required"`
}

func (i *Request) MapIntoEmbedding() (entity.Embedding, error) {
	var embedding entity.Embedding = entity.Embedding{
		UUID:        uuid.New().String(),
		Category:    i.Category,
		Granularity: i.Granularity,
		Content:     i.Content,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := embedding.Validate(); err != nil {
		return entity.Embedding{}, err
	}
	return embedding, nil
}
