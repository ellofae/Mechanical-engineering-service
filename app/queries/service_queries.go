package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/ellofae/Mechanical-engineering-service/app/models"
)

type ServiceQueries struct {
	*sqlx.DB
}

func (s *ServiceQueries) GetServices() ([]models.Service, error) {
	services := []models.Service{}

	query := `SELECT * FROM services`

	err := s.Select(&services, query)
	if err != nil {
		return services, err
	}

	return services, nil
}

func (sq *ServiceQueries) CreateService(s *models.Service) error {
	query := `INSERT INTO services VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx := sq.MustBegin()
	_, err := tx.Exec(query, s.ID, s.Created_at, s.Updated_at, s.Service_title, s.Service_description, s.Service_status, s.Detail_model)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (sq *ServiceQueries) UpdateService(id uuid.UUID, s *models.Service) error {
	query := `UPDATE services SET updated_at = $2, service_title = $3, service_description = $4, service_status = $5, detail_model = $6 WHERE id = $1`

	tx := sq.MustBegin()
	_, err := tx.Exec(query, id, s.Updated_at, s.Service_title, s.Service_description, s.Service_status, s.Detail_model)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (sq *ServiceQueries) DeleteService(id uuid.UUID) error {
	query := `DELETE FROM services WHERE id = $1`

	tx := sq.MustBegin()
	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}