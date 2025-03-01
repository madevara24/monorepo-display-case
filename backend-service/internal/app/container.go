package app

import (
	"backend-service/config"
	questionRepository "backend-service/internal/app/domain/question/repository"
	"backend-service/internal/app/domain/question/usecase/ask"
	"backend-service/internal/app/domain/system/usecase/healthcheck"
	"backend-service/internal/pkg/datasource"
	"backend-service/internal/pkg/integration/openai"
)

type Container struct {
	// Services
	HealthCheckService *healthcheck.HealthCheckService
	AskService         *ask.AskService

	// External Services
	OpenAIClient *openai.Client
}

func NewContainer(datasource *datasource.DataSource) *Container {
	openAIClient := openai.NewClient(config.Get().OpenAIAPIKey)
	questionRepo := questionRepository.NewQuestionRepository(datasource.Postgres)
	return &Container{
		// Services
		HealthCheckService: healthcheck.NewHealthCheckService(),
		AskService:         ask.NewAskService(openAIClient, questionRepo),

		// External Services
		OpenAIClient: openAIClient,
	}
}
