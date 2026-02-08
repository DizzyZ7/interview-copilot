package questions

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) Create(q *Question) error {
	return r.DB.QueryRow(context.Background(),
		`INSERT INTO questions(title, body, difficulty, tags)
		 VALUES($1,$2,$3,$4) RETURNING id`,
		q.Title, q.Body, q.Difficulty, q.Tags,
	).Scan(&q.ID)
}

func (r *Repository) List(tag string, difficulty int) ([]Question, error) {
	q := `SELECT id, title, body, difficulty, tags FROM questions WHERE 1=1`
	args := []any{}
	i := 1

	if tag != "" {
		q += " AND $1 = ANY(tags)"
		args = append(args, tag)
		i++
	}
	if difficulty > 0 {
		q += " AND difficulty = $" + string(rune(i+'0'))
		args = append(args, difficulty)
	}

	rows, err := r.DB.Query(context.Background(), q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Title, &q.Body, &q.Difficulty, &q.Tags)
		if err != nil {
			return nil, err
		}
		res = append(res, q)
	}
	return res, nil
}

func (r *Repository) Get(id int) (*Question, error) {
	var q Question
	err := r.DB.QueryRow(context.Background(),
		`SELECT id, title, body, difficulty, tags FROM questions WHERE id=$1`, id).
		Scan(&q.ID, &q.Title, &q.Body, &q.Difficulty, &q.Tags)
	return &q, err
}

func (r *Repository) Update(id int, q *Question) error {
	_, err := r.DB.Exec(context.Background(),
		`UPDATE questions SET title=$1, body=$2, difficulty=$3, tags=$4 WHERE id=$5`,
		q.Title, q.Body, q.Difficulty, q.Tags, id)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.DB.Exec(context.Background(),
		`DELETE FROM questions WHERE id=$1`, id)
	return err
}

func (r *Repository) Random(limit int) ([]Question, error) {
	rows, err := r.DB.Query(context.Background(),
		`SELECT id, title, body, difficulty, tags
		 FROM questions ORDER BY random() LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Question
	for rows.Next() {
		var q Question
		rows.Scan(&q.ID, &q.Title, &q.Body, &q.Difficulty, &q.Tags)
		res = append(res, q)
	}
	return res, nil
}
