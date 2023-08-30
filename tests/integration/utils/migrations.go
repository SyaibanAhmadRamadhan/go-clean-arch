package utils

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func StartMigration(url string, db *sql.DB) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Err(err).Msgf("could not start db: %s", err)
		os.Exit(1)
	}
	log.Info().Msgf("start migrate")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Err(err).Msgf("could not init driver: %s", err)
		os.Exit(1)
	}

	migrates, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		log.Err(err).Msgf("could not apply the migration: %s", err)
		os.Exit(1)
	}
	migrates.Up()

}
