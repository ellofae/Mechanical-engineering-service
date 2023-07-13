package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID                    uuid.UUID  `db:"id" json:"id" validate:"required,uuid"`
	Created_at            time.Time  `db:"created_at" json:"created_at"`
	Updated_at            time.Time  `db:"updated_at" json:"updated_at"`
	Vehicle_price         string     `db:"vehicle_price" json:"vehicle_price" valudate:"required,lte=128"`
	Cetegory              string     `db:"category" json:"category" valudate:"required,lte=255"`
	Title                 string     `db:"title" json:"title" valudate:"required,lte=255"`
	Vehicle_status        string     `db:"vehicle_status" json:"vehicle_status" valudate:"required,lte=128"`
	Model                 string     `db:"model" json:"model" valudate:"required,lte=255"`
	Model_description     string     `db:"model_description" json:"model_description" valudate:"required,lte=255"`
	Model_characteristics ModelChars `db:"model_characteristics" json:"model_characteristics validate:"required,dive"`
}

type ModelChars struct {
	Year        string  `db:"year" json:"year"`
	Mileage     float32 `db:"mileage" json:"mileage" validate:"required,gte=0"`
	Engine      string  `db:"engine" json:"engine" validate:"required,lte=255"`
	Engine_spec string  `db:"engine_spec" json:"engine_spec" validate:"required,lte=255"`
	Suspensions string  `db:"suspensions" json:"suspensions" validate:"required,lte=255"`
	Bodykit     string  `db:"bodykit" json:"bodykit" validate:"required,lte=255"`
	Remarks     string  `db:"remarks" json:"remarks" validate:"required,lte=255"`
}

func (mc ModelChars) Value() (driver.Value, error) {
	return json.Marshal(mc)
}

func (mc *ModelChars) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &mc)
}
