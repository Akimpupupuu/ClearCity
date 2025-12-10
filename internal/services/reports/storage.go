package reports

import (
	"github.com/Akimpupupuu/ClearCity/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{DB: db}
}

func (storage *Storage) CreateReport(types.Report) error {
	return nil
}

func (storage *Storage) GetReports() error {
	return nil
}
