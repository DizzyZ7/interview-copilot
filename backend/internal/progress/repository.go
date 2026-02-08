package progress

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) Record(userID, questionID int, correct bool) error {
	_, err := r.DB.Exec(context.Background(),
		`INSERT INTO progress(user_id, question_id, correct)
		 VALUES($1,$2,$3)`, userID, questionID, correct)
	return err
}

func (r *Repository) Stats(userID int) (total int, correct int, err error) {
	err = r.DB.QueryRow(context.Background(),
		`SELECT COUNT(*), COUNT(*) FILTER (WHERE correct=true)
		 FROM progress WHERE user_id=$1`, userID).
		Scan(&total, &correct)
	return
}
