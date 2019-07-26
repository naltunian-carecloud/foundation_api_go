package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CareCloud/foundation_api_go/middleware"
	"github.com/CareCloud/foundation_api_go/routes/v1/bundle"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

type server struct {
	router *mux.Router
	srv    *http.Server
}

func main() {

	s := server{
		router: mux.NewRouter(),
	}

	s.routes()
	port := os.Getenv("CONFIG_HTTP_PORT")

	if port == "" {
		port = "4011"
	}
	log.Println("Using Port " + port)

	s.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        s.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.srv.ListenAndServe()

	if err != nil {
		log.Fatal(errors.New("Server failed to start! 💥"))
	}

}

func (s *server) routes() {

	s.router.Use(middleware.Logger)

	s.router.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	}).Methods("GET")

	v1 := s.router.PathPrefix("/api/v1").Subrouter()
	bundle.Routes(v1)
}
