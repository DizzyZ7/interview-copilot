package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func (r *UserRepo) Create(email, hash string) error {
	_, err := r.DB.Exec(context.Background(),
		"INSERT INTO users(email, password_hash) VALUES($1,$2)", email, hash)
	return err
}

func (r *UserRepo) GetByEmail(email string) (int, string, error) {
	var id int
	var hash string
	err := r.DB.QueryRow(context.Background(),
		"SELECT id, password_hash FROM users WHERE email=$1", email).
		Scan(&id, &hash)
	return id, hash, err
}
