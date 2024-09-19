package server

import (
	"fmt"
	"i-am-here/app/internal/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

func IAHServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	IAHServer := &Server{
		port: port,
		db:   database.New(),
	}

	mux := IAHServer.RegisterRoutes()
	// Set up CORS handling
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // Allow requests from this origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allow these HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Allow these headers
	)(mux)
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", IAHServer.port),
		Handler:      corsHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
