package drivens

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func CreatingConnection() (*pgxpool.Pool, error) {
	err := godotenv.Load("../go.env")
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		errorMessage := fmt.Sprintf("Unable to create connection pool: %v\n", err)
		return nil, errors.New(errorMessage)
	}

	return dbpool, nil
}

type dbHandler struct {
	connection *pgxpool.Pool
}

func (*d dbdbHandler) 