package helper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "friends-management"
)

func ConnectDb() *sql.DB {

	dataInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open db connection
	db, err := sql.Open("postgres", dataInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func ConnectDbFailed() *sql.DB {

	dataInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "postgres", "11", "friends-management")

	// Open db connection
	db, err := sql.Open("postgres", dataInfo)
	if err != nil {
		panic(err)
	}

	return db
}
