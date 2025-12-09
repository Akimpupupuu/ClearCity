package reports

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Storage *Storage
}

func NewHandler(storage *Storage) *Handler {
	return &Handler{Storage: storage}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_report", handler.CreateReport).Methods("POST")
	router.HandleFunc("/get_reports", handler.GetReports).Methods("GET")
}

func (handler *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetReports(w http.ResponseWriter, r *http.Request) {
	//get json request and parse it into struct-payload
	//work with db
	//write json response
}
