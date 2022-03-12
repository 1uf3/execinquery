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

	_ = db.QueryRowContext(context.Background(), "SELECT * FROM test WHERE test=?", s)

	_ = db.QueryRowContext(context.Background(), "DELETE * FROM test WHERE test=?", s) // want "QueryRowContext\\(\\) not recommended execute `DELETE` query"
	_ = db.QueryRowContext(context.Background(), "UPDATE * FROM test WHERE test=?", s) // want "QueryRowContext\\(\\) not recommended execute `UPDATE` query"

	_, _ = db.Query("UPDATE * FROM test WHERE test=?", s)                              // want "Query\\(\\) not recommended execute `UPDATE` query"
	_, _ = db.QueryContext(context.Background(), "UPDATE * FROM test WHERE test=?", s) // want "QueryContext\\(\\) not recommended execute `UPDATE` query"
	_ = db.QueryRow("UPDATE * FROM test WHERE test=?", s)                              // want "QueryRow\\(\\) not recommended execute `UPDATE` query"

	query := "UPDATE * FROM test where test=?"
	_, _ = db.Query(query, s) //want "Query\\(\\) not recommended execute `UPDATE` query"
}
