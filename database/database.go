package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}

func MigrateDB() {
	const createTableQuery = `
    CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sender TEXT NOT NULL,
        receiver TEXT NOT NULL,
        content TEXT NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migration completed")
}
