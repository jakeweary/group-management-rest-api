package database

import (
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	sqlx *sqlx.DB
}

func Connect(url string) (*Database, error) {
	slog.Debug("connecting to database")
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, err
	}

	driver, err := pgx.WithInstance(db.DB, &pgx.Config{})
	if err != nil {
		return nil, err
	}

	slog.Debug("initializing database migrator")
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}

	// slog.Debug("migrating down")
	// if err := migrator.Down(); err != nil {
	// 	slog.Warn("failed to migrate down", "err", err)
	// }

	slog.Debug("migrating up")
	if err := migrator.Up(); err != nil {
		slog.Warn("failed to migrate up", "err", err)
	}

	return &Database{db}, nil
}

func (db *Database) Close() {
	slog.Debug("closing database")
	db.sqlx.Close()
}
