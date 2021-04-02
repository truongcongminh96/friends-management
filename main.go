package main

import (
	"fmt"
	"github.com/friends-management/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main()  {
	PORT := ":8081"

	err := godotenv.Load()
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

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
