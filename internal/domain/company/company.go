package company

import (
	"database/sql"
	"github.com/google/uuid"
)

type Entity struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	Phone       string
	Description sql.NullString
	AccountType string
	CreatedAt   string
	UpdatedAt   string
}
