package main

import (
	"log"
	"strings"

	"github.com/nollidnosnhoj/simpbb/internal/database"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func GetDbCommand(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name: "migrate",
		Usage: "Database migration commands",
		Subcommands: []*cli.Command{
			{
				Name: "init",
				Usage: "Create migration tables",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					return migrator.Init(ctx.Context)
				},
			},
			{
				Name: "up",
				Usage: "Run all pending migrations",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					group, err := migrator.Migrate(ctx.Context)
					if err != nil {
						return err
					}
					if group.ID == 0 {
						log.Println("No migrations to run")
						return nil
					}
					log.Printf("Migrated to version %s", group.ID)
					return nil
				},
			},
			{
				Name: "down",
				Usage: "Rollback the last migration",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					group, err := migrator.Rollback(ctx.Context)
					if err != nil {
						return err
					}
					if group.ID == 0 {
						log.Println("No migrations to rollback")
						return nil
					}
					log.Printf("Rolled back to version %s", group.ID)
					return nil
				},
			},
			{
				Name: "lock",
				Usage: "Lock migration",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					return migrator.Lock(ctx.Context)
				},
			},
			{
				Name: "unlock",
				Usage: "Unlock migration",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					return migrator.Unlock(ctx.Context)
				},
			},
			{
				Name: "status",
				Usage: "Show migration status",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					ms, err := migrator.MigrationsWithStatus(ctx.Context)
					if err != nil {
						return err
					}
					log.Printf("migrations: %s\n", ms)
					log.Printf("unapplied migrations: %s\n", ms.Unapplied())
					log.Printf("last migration group: %s\n", ms.LastGroup())

					return nil
				},
			},
			{
				Name: "create",
				Usage: "Create a new migration",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)
					name := strings.Join(ctx.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(ctx.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						log.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(ctx *cli.Context) error {
					db := database.NewDb()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations)

					group, err := migrator.Migrate(ctx.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}

					if group.ID == 0 {
						log.Printf("there are no new migrations to mark as applied\n")
						return nil
					}

					log.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}
