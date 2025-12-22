package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Akimpupupuu/ClearCity/internal/config"
	"github.com/Akimpupupuu/ClearCity/internal/services/reports"
	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()

	reportStorage := reports.NewStorage(server.DB)
	reportHandler := reports.NewHandler(reportStorage)

	router.Route("/v1", func(rv1 chi.Router) {
		rv1.Route("/reports", func(reportsRouter chi.Router) {
			reportsRouter.Post("/create", reportHandler.CreateReport)
		})
	})

	log.Println("Listenning on", server.Address)
	return http.ListenAndServe(fmt.Sprintf(":%d", server.Address), router)
}
