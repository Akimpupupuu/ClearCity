package types

import "time"

type ReportStorage interface {
	CreateReport(Report) error
	GetReports() error
}

type UserStorage interface {
}

type ReportPayload struct {
	Address     string
	Description string
}

type Report struct {
	Id          int
	UserId      int
	Address     string
	Description string
	CreatedAt   time.Time
}

type UserPayload struct {
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type User struct {
	Id          int
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}
