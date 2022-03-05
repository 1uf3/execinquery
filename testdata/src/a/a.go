package a

import (
	"context"
	"database/sql"
	"log"
	"testing"
)

func setup() *sql.DB {
	db, err := sql.Open("mysql", "test:test@tcp(test:3306)/test")
	if err != nil {
		log.Fatal("Database Connect error: ", err)
	}
	return db
}

func f(t *testing.T) {
	db := setup()
	defer db.Close()

	_ = db.QueryRowContext(context.Background(), "SELECT * FROM comments WHERE user=?", "alice")
	// want
	_ = db.QueryRowContext(context.Background(), "DELETE * FROM comments WHERE user=?", "alice")
	// want "QueryRowContext() can not use "DELETE" query"
}
