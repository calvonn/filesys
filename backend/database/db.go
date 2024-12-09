package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB_Conn *sql.DB

func InitDB() {
	var err error
	DB_Conn, err = sql.Open("postgres", "user=calvon password=123456 dbname=filesys sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := DB_Conn.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
