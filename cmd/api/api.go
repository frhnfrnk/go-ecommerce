package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/frhnfrnk/go-ecommerce/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.Newhandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	
	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}