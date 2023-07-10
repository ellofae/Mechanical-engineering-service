package database

import (
	"github.com/ellofae/Mechanical-engineering-service/app/queries"
)

type Queries struct {
	*queries.ServiceQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{ServiceQueries: &queries.ServiceQueries{DB: db}}, nil
}