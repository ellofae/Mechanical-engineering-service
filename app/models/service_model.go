package models

import (
	"time"
	"github.com/google/uuid"
)

type Service struct {
	ID uuid.UUID `db:"id" json:"id" validate="required,uuid"`
    Created_at time.Time `db:"created_at" json:"created_at"`
    Updated_at time.Time `db:"updated_at" json:"updated_at"`
    Service_title string `db:"service_title" json:"service_title" validate="required,lte=255"`
    Service_description string `db:"service_description" json:"service_description" validate="required,lte=255"`
    Service_status string `db:"service_status" json:"service_status" validate="required,lte=128"`
    Detail_model string `db:"detail_model" json:"detail_model" validate="required,lte=255"`
}