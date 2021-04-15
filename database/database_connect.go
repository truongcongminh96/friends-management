package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	HOST = "127.0.0.1"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func ConnectDB() (Database, error) {
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		database = os.Getenv("POSTGRES_DB")
	)

	db := Database{}

	dataInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username, password, HOST, PORT, database)

	conn, err := sql.Open("postgres", dataInfo)

	if err != nil {
		return db, err
	}

	db.Conn = conn

	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Successfully connected!")
	return db, nil
}
