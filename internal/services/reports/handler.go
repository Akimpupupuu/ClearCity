package reports

import (
	"fmt"
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
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	report, err := handler.Storage.CreateReport(types.Report{
		Title:       payload.Title,
		Address:     payload.Address,
		Description: payload.Description,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("this is our report: %s\n", report.Description)

	utils.WriteJSON(w, http.StatusCreated, nil)
}
