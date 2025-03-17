package store

import (
	"backend-service/internal/app/domain/knowledge/entity"
	"backend-service/internal/app/domain/knowledge/repository"
	"backend-service/internal/pkg/integration/openai"
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	openAIClient  *openai.Client
	knowledgeRepo *repository.KnowledgeRepository
}

func NewStoreUsecase(openAIClient *openai.Client, knowledgeRepo *repository.KnowledgeRepository) *StoreUsecase {
	return &StoreUsecase{
		openAIClient:  openAIClient,
		knowledgeRepo: knowledgeRepo,
	}
}

func (i *StoreUsecase) Execute(ctx context.Context, req Request) error {
	knowledge, err := req.MapIntoKnowledge()
	if err != nil {
		return err
	}

	// Create embedding for the question
	vector, err := i.openAIClient.CreateEmbedding(ctx, req.Content)
	if err != nil {
		return err
	}

	knowledge.Embedding = vector

	return i.knowledgeRepo.StoreKnowledge(ctx, knowledge)
}

func (i *StoreUsecase) ExecuteV2(ctx context.Context, req RequestV2) error {
	knowledge, err := req.MapIntoKnowledgeV2()
	if err != nil {
		return err
	}

	// Create embedding for the question
	vector, err := i.openAIClient.CreateEmbeddingV2(ctx, req.Content)
	if err != nil {
		return err
	}

	knowledge.Embedding = vector

	return i.knowledgeRepo.StoreKnowledge(ctx, knowledge)
}

type Request struct {
	Category    string `json:"category" validate:"required"`
	Granularity string `json:"granularity" validate:"required"`
	Content     string `json:"content" validate:"required"`
}

type RequestV2 struct {
	entity.KnowledgeRaw
}

func (r *Request) MapIntoKnowledge() (entity.Knowledge, error) {
	var knowledge entity.Knowledge = entity.Knowledge{
		UUID:        uuid.New().String(),
		Category:    r.Category,
		Granularity: r.Granularity,
		Content:     r.Content,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := knowledge.Validate(); err != nil {
		return entity.Knowledge{}, err
	}
	return knowledge, nil
}

func (r *RequestV2) MapIntoKnowledgeV2() (entity.Knowledge, error) {

	metadataJSON, err := json.Marshal(r.Metadata)
	if err != nil {
		return entity.Knowledge{}, err
	}
	category := string(metadataJSON)

	knowledge := entity.Knowledge{
		UUID:        uuid.New().String(),
		Category:    category,
		Granularity: r.Metadata.Granularity,
		Content:     r.Content,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := knowledge.Validate(); err != nil {
		return entity.Knowledge{}, err
	}
	return knowledge, nil
}
