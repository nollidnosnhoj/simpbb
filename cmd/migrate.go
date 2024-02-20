package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

var MigrateCommand = &cli.Command{
	Name: "migrate",
	Usage: "migrate the database",
	Subcommands: []*cli.Command{
		{
			Name: "up",
			Usage: "migrate the database up",
			Action: upMigration,
		},
		{
			Name: "down",
			Usage: "migrate the database down",
			Action: downMigration,
		},
	},
}

func upMigration(c *cli.Context) error {
	m, err := getMigrator()
	if err != nil {
		return err
	}
	defer func () {
		err1, err2 := m.Close()
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err2)
		}
	}()
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func downMigration(c *cli.Context) error {
	m, err := getMigrator()
	if err != nil {
		return err
	}
	defer func () {
		err1, err2 := m.Close()
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err2)
		}
	}()
	if err := m.Down(); err != nil {
		return err
	}
	return nil
}

func getMigrator() (*migrate.Migrate, error) {
	db, err := sql.Open("sqlite3", "file:simpbb.db")
	if err != nil {
		return nil, err
	}
	driver, driverErr := sqlite3.WithInstance(db, &sqlite3.Config{})
	if driverErr != nil {
		return nil, driverErr
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}