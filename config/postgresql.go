package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDatabase() (*sql.DB, error) {
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USR := os.Getenv("POSTGRES_USR")
	POSTGRES_PWD := os.Getenv("POSTGRES_PWD")

	POSTGRES_URL := fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		POSTGRES_USR, POSTGRES_PWD, POSTGRES_HOST, POSTGRES_DB,
	)

	client, err := sql.Open("postgres", POSTGRES_URL)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
