package app

import (
	"backend-service/internal/app/domain/question/usecase/ask"
	"backend-service/internal/app/domain/system/usecase/healthcheck"
	"backend-service/internal/pkg/datasource"
)

type Container struct {
	// Services
	HealthCheckService *healthcheck.HealthCheckService
	AskService         *ask.AskService
}

func NewContainer(datasource *datasource.DataSource) *Container {
	return &Container{
		HealthCheckService: healthcheck.NewHealthCheckService(),
		AskService:         ask.NewAskService(),
	}
}
