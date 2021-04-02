package main

import (
	"fmt"
	"github.com/friends-management/database"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

func main()  {
	address := ":8081"
	network, err := net.Listen("tcp", address)

	fmt.Println(network.Addr(), err)

	if err != nil {
		// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
		// /src/log/log.go:208
		log.Fatalf("Error: %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	db, err := database.ConnectDB(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer db.Conn.Close()
}
