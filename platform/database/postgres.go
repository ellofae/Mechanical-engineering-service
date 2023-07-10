package database

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"	

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v4/stdlib"
)

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