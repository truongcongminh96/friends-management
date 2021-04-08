package main

import (
	"context"
	"github.com/friends-management/database"
	"github.com/friends-management/handlers"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	PORT := ":8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	db, err := database.ConnectDB(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Conn.Close()

	httpHandler := handlers.NewHandler(db)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", PORT)
	ch := make(chan os.Signal, 1)

	// https://tour.golang.org/concurrency/1
	// https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
	// The API server is started on a separate goroutine and keeps running until it receives a SIGINT or SIGTERM signal after which it calls the Stop function to clean up and shut down the server.
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Go channels
	//done := make(chan bool)
	//go hello(done)
	//<-done
	//fmt.Println("main function")
	<-ch
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed: %v\n", err)
		os.Exit(1)
	}
}
