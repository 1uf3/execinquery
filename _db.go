package main

import (
	"context"
	"database/sql"
	"log"
)

func main() {
	db, err = sql.Open("mysql", "test:test@tcp(test:3306)/test")
	if err != nil {
		log.Fatal("Database Connect error: ", err)
	}
	defer db.Close()

	var user string
	result := db.QueryRowContext(context.Background(), "SELECT * FROM comments WHERE user=?", "alice")
	if err := result.Scan(&user); err != nil {
		return err
	}

	result := db.QueryRowContext(context.Background(), "DELETE * FROM comments WHERE user=?", "alice")
	if err := result.Scan(&user); err != nil {
		return err
	}
	return nil
}
