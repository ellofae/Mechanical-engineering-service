package queries

import (
	"github.com/ellofae/Mechanical-engineering-service/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VehicleQueries struct {
	*sqlx.DB
}

func (v *VehicleQueries) GetVehicles() ([]models.Vehicle, error) {
	vehicles := []models.Vehicle{}

	query := `SELECT * FROM vehicles`

	err := v.Select(&vehicles, query)
	if err != nil {
		return vehicles, err
	}

	return vehicles, nil
}

func (v *VehicleQueries) GetVehicle(id uuid.UUID) (models.Vehicle, error) {
	vehicle := models.Vehicle{}

	query := `SELECT * FROM vehicles WHERE id = $1`

	err := v.Get(&vehicle, query, id)
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}

func (v *VehicleQueries) CreateVehicle(vm *models.Vehicle) error {
	query := `INSERT INTO vehicles VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	tx := v.MustBegin()
	_, err := tx.Exec(query, vm.ID, vm.Created_at, vm.Updated_at, vm.Vehicle_price, vm.Cetegory, vm.Title, vm.Vehicle_status, vm.Model, vm.Model_description, vm.Model_characteristics)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (v *VehicleQueries) UpdateVehicle(id uuid.UUID, vm *models.Vehicle) error {
	query := `UPDATE vehicles SET updated_at = $2, vehicle_price = $3, category = $4, title = $5, vehicle_status = $6, model = $7, model_description = $8, model_characteristics = &9 WHERE id = $1`

	tx := v.MustBegin()
	_, err := tx.Exec(query, vm.ID, vm.Updated_at, vm.Vehicle_price, vm.Cetegory, vm.Title, vm.Vehicle_status, vm.Model, vm.Model_description, vm.Model_characteristics)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (v *VehicleQueries) DeleteVehicle(id uuid.UUID) error {
	query := `DELETE FROM vehicles WHERE id = $1`

	tx := v.MustBegin()
	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
