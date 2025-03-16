package app

import (
	"backend-service/config"
	knowledgeRepository "backend-service/internal/app/domain/knowledge/repository"
	"backend-service/internal/app/domain/knowledge/usecase/store"
	questionRepository "backend-service/internal/app/domain/question/repository"
	"backend-service/internal/app/domain/question/usecase/ask"
	"backend-service/internal/app/domain/system/usecase/healthcheck"
	"backend-service/internal/pkg/datasource"
	"backend-service/internal/pkg/integration/openai"
)

type Container struct {
	// Services
	HealthCheckUsecase *healthcheck.HealthCheckUsecase
	AskUsecase         *ask.AskUsecase
	StoreUsecase       *store.StoreUsecase

	// External Services
	OpenAIClient *openai.Client
}

func NewContainer(datasource *datasource.DataSource) *Container {
	openAIClient := openai.NewClient(config.Get().OpenAIAPIKey)
	questionRepo := questionRepository.NewQuestionRepository(datasource.Postgres)
	knowledgeRepo := knowledgeRepository.NewKnowledgeRepository(datasource.Postgres)
	return &Container{
		// Services
		HealthCheckUsecase: healthcheck.NewHealthCheckService(datasource.Postgres),
		AskUsecase:         ask.NewAskUsecase(datasource.Postgres, openAIClient, questionRepo, knowledgeRepo),
		StoreUsecase:       store.NewStoreUsecase(openAIClient, knowledgeRepo),

		// External Services
		OpenAIClient: openAIClient,
	}
}
