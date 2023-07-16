package database

import (
	"fmt"
	"os"

	"github.com/ellofae/Mechanical-engineering-service/app/queries"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.ServiceQueries
	*queries.VehicleQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		ServiceQueries: &queries.ServiceQueries{DB: db},
		VehicleQueries: &queries.VehicleQueries{DB: db},
	}, nil
}

func PostgreSQLConnection() (*sqlx.DB, error) {
	godotenv.Load(".env")

	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database. %w", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database. %w", err)
	}

	return db, nil
}
