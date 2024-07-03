package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/books-management?sslmode=disable")
	if err != nil {
		fmt.Println("fail");
		panic(err)
	}
	DB = db;
	err = createBooksTable();
	if err != nil {
		// fmt.Println("fail");
		panic(err);
	}
}

func createBooksTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title TEXT,
		author TEXT,
		year INT
	);
	`
	_, err := DB.Exec(createTableQuery)
	return err
}
