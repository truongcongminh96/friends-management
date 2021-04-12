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

type envVars struct {
	Port       string `env:"PORT"`
	DbUser     string `env:"POSTGRES_USER"`
	DbPassword string `env:"POSTGRES_PASSWORD"`
	DbName     string `env:"POSTGRES_DB"`
}

var envConfig envVars

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envConfig.Port = os.Getenv("PORT")
	envConfig.DbUser = os.Getenv("POSTGRES_USER")
	envConfig.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	envConfig.DbName = os.Getenv("POSTGRES_DB")
}

func main() {
	db, err := database.ConnectDB(envConfig.DbUser, envConfig.DbPassword, envConfig.DbName)

	listener, err := net.Listen("tcp", ":"+envConfig.Port)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	defer db.Conn.Close()

	httpHandler := routes.NewHandler(db)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", envConfig.Port)
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
