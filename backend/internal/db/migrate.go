package db

import (
	"context"
	"os"
	"path/filepath"
	"sort"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(db *pgxpool.Pool) {
	files, _ := filepath.Glob("migrations/*.sql")
	sort.Strings(files)

	for _, f := range files {
		sql, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(context.Background(), string(sql))
		if err != nil {
			panic(err)
		}
	}
}
