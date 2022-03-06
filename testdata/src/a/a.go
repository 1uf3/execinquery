package a

import (
	"context"
	"database/sql"
	"testing"
)

func setup() *sql.DB {
	db, _ := sql.Open("mysql", "test:test@tcp(test:3306)/test")
	return db
}

func f(t *testing.T) {
	db := setup()
	defer db.Close()

	s := "alice"

	_ = db.QueryRowContext(context.Background(), "SELECT * FROM comments WHERE user=?", s)

	_ = db.QueryRowContext(context.Background(), "DELETE * FROM comments WHERE user=?", s) // want "QueryRowContext\\(\\) can not use `DELETE` query"
	_ = db.QueryRowContext(context.Background(), "UPDATE * FROM comments WHERE user=?", s) // want "QueryRowContext\\(\\) can not use `UPDATE` query"
}
