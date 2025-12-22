package reports

import (
	"context"

	"github.com/Akimpupupuu/ClearCity/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{DB: db}
}

func (storage *Storage) CreateReport(report types.Report) (*types.Report, error) {
	query := `INSERT INTO report (title, address, description) VALUES ($1, $2, $3) RETURNING id, created_at`

	if err := storage.DB.QueryRow(context.Background(), query, report.Title, report.Address, report.Description).Scan(
		&report.Id,
		&report.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &report, nil
}
