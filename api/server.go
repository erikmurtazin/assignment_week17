package api

import (
	"assignment_week17/db"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Memory map[string]string
	Db     *db.Mongodb
}

func NewServer(db *db.Mongodb) *Server {
	return &Server{
		Memory: map[string]string{},
		Db:     db,
	}
}

func (s *Server) Run() {
	router := http.NewServeMux()
	router.HandleFunc("/mongo", s.HandleMongoRequest)
	router.HandleFunc("/in-memory", s.HandleInMemoryRequest)
	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Println("Server listening on port", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
