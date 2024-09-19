package server

import (
	"encoding/json"
	"i-am-here/app/internal/models"
	"log"
	"net/http"
	"time"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/users", s.GetUsersHandler)
	mux.HandleFunc("/test", s.TestHandler)
	mux.HandleFunc("/user-create", s.CreateUserHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Test request...")
	testData, err := s.db.GetTestData()
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched data: %v", testData)
	jsonResp, err := json.Marshal(testData)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		http.Error(w, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling GetUsers request...")
	users, err := s.db.GetUsers()
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched users: %v", users)
	jsonResp, err := json.Marshal(users)
	if err != nil {
		log.Printf("Failed to marshal users: %v", err)
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := s.db.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
}
