package users

import (
	"net/http"

	"github.com/Akimpupupuu/ClearCity/internal/types"
)

type Handler struct {
	Storage types.UserStorage
}

func NewHandler(storage types.UserStorage) *Handler {
	return &Handler{Storage: storage}
}

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
