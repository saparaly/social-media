package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	var queries []string = []string{users, session, post, postLikes, postDislikes, postDislikes, comment, commentLikes, commentDislikes}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Println("tables not created")
			return err
		}
	}

	return nil
}
