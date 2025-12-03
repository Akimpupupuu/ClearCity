package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Akimpupupuu/ClearCity/internal/config"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HttpServer struct {
	Address int
	DB      *pgxpool.Pool
}

func NewServer(cfg *config.Config, db *pgxpool.Pool) *HttpServer {
	return &HttpServer{
		Address: cfg.HTTPPort,
		DB:      db,
	}
}

func (server *HttpServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Listenning on", server.Address)
	return http.ListenAndServe(fmt.Sprintf(":%d", server.Address), subrouter)
}
