package types

import "time"

type ReportStorage interface {
	CreateReport(Report) (*Report, error)
}

type ReportPayload struct {
	Title       string `json:"title"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Report struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
