package users

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

func (handler *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/delete", handler.Delete).Methods("DELETE")
}

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
