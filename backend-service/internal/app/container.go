package app

import (
	"backend-service/config"
	questionRepository "backend-service/internal/app/domain/question/repository"
	"backend-service/internal/app/domain/question/usecase/ask"
	"backend-service/internal/app/domain/question/usecase/store"
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
	return &Container{
		// Services
		HealthCheckUsecase: healthcheck.NewHealthCheckService(datasource.Postgres),
		AskUsecase:         ask.NewAskUsecase(openAIClient, questionRepo),
		StoreUsecase:       store.NewStoreUsecase(openAIClient, questionRepo),

		// External Services
		OpenAIClient: openAIClient,
	}
}
