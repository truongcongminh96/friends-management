package main

import (
	"context"
	"github.com/friends-management/database"
	"github.com/friends-management/routes"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	PORT := os.Getenv("PORT")
	db, err := database.ConnectDB()

	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	defer db.Conn.Close()

	httpHandler := routes.NewHandler(db)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		_ = server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", PORT)
	ch := make(chan os.Signal, 1)

	/* The API server is started on a separate goroutine and keeps running until it receives a SIGINT or SIGTERM
	   signal after which it calls the Stop function to clean up and shut down the server. */
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
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
