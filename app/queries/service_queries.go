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
	tx.MustExec(query, s.ID, s.Created_at, s.Updated_at, s.Service_title, s.Service_description, s.Service_status, s.Detail_model)
	tx.Commit()

	return nil
}