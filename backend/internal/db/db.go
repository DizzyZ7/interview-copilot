package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(url string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to PostgreSQL")
	return pool
}
