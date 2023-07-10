package models

import (
	"time"
	"github.com/google/uuid"
)

type Service struct {
	ID uuid.UUID `db:"id" json:"id"`
    Created_at time.Time `db:"created_at" json:"created_at"`
    Updated_at time.Time `db:"updated_at" json:"updated_at"`
    Service_title string `db:"title" json:"title"`
    Service_description string `db:"description" json:"description"`
    Service_status string `db:"status" json:"status"`
    Detail_model string `db:"detail_model" json:"detail_model"`
}