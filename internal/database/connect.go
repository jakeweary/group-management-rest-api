package database

import (
	"log"
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

func Connect(url string) Database {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatalln(err)
	}

	driver, err := pgx.WithInstance(db.DB, &pgx.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalln(err)
	}

	// slog.Debug("migrating down")
	// if err := migrator.Down(); err != nil {
	// 	slog.Info("failed to migrade down", "err", err)
	// }

	slog.Debug("migrating up")
	if err := migrator.Up(); err != nil {
		slog.Info("failed to migrade up", "err", err)
	}

	return Database{db}
}

func (db *Database) Close() {
	db.sqlx.Close()
}
