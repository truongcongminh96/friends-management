package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	HOST = "127.0.0.1"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func ConnectDB(username, password, database string) (Database, error) {
	db := Database{}

	dataInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dataInfo)
	fmt.Println(conn, err)

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
