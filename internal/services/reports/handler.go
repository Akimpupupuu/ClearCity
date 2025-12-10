package reports

import (
	"net/http"

	"github.com/Akimpupupuu/ClearCity/internal/types"
	"github.com/Akimpupupuu/ClearCity/internal/utils"
)

type Handler struct {
	Storage types.ReportStorage
}

func NewHandler(storage types.ReportStorage) *Handler {
	return &Handler{Storage: storage}
}

func (handler *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var payload types.ReportPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := handler.Storage.CreateReport(types.Report{
		Address:     payload.Address,
		Description: payload.Description,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (handler *Handler) GetReports(w http.ResponseWriter, r *http.Request) {

}
