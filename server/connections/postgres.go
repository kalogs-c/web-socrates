package connections

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/kalogs-c/web-socrates/internal/sqlc"
)

type PostgresConnection struct {
	db *sql.DB
	Q  *sqlc.Queries
}

func NewPostgresConnection() (*PostgresConnection, error) {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	db := sqlc.New(conn)
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://sql/migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return &PostgresConnection{conn, db}, nil
}

func (p *PostgresConnection) Close() error {
	return p.db.Close()
}
