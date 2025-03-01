package datasource

import (
	"backend-service/config"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/madevara24/go-common/database"
)

type DataSource struct {
	Postgres *sqlx.DB
}

func NewDataSource() *DataSource {
	postgresClient := database.NewConfiguration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Get().DBUsername,
		config.Get().DBPassword,
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBName,
	), "backend-service-sqlx")

	postgresDB, err := sqlx.Connect("postgres", postgresClient.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	postgresDB.SetMaxIdleConns(config.Get().DBMaxIdleConn)
	postgresDB.SetMaxOpenConns(config.Get().DBMaxConn)
	postgresDB.SetConnMaxLifetime(time.Duration(config.Get().DBMaxTTLConn) * time.Second)

	return &DataSource{
		Postgres: postgresDB,
	}
}
