package healthcheck

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type HealthCheckUsecase struct {
	db *sqlx.DB
}

func NewHealthCheckService(db *sqlx.DB) *HealthCheckUsecase {
	return &HealthCheckUsecase{
		db: db,
	}
}

func (i *HealthCheckUsecase) Execute(ctx context.Context) (Response, error) {
	status := Status{
		Postgresql: true,
	}

	err := i.db.PingContext(ctx)
	if err != nil {
		return Response{
			Message: "UNHEALTHY",
			Status: Status{
				Postgresql: false,
			},
		}, err
	}

	return Response{
		Message: "OK",
		Status:  status,
	}, nil
}

type Response struct {
	Message string `json:"message"`
	Status  Status `json:"status"`
}

type Status struct {
	Postgresql bool `json:"postgresql"`
}
