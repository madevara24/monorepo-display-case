package healthcheck

import (
	"context"
)

type HealthCheckService struct {
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (i *HealthCheckService) Execute(ctx context.Context) (Response, error) {
	return Response{
		Status: "OK",
	}, nil
}

type Response struct {
	Status string `json:"status"`
}
