package bootstrap

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
)

var (
	migrationup = &cobra.Command{
		Use:   "migrate:up",
		Short: "Execute Database Migration Up",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New("file://files/migration", cfg.GetString("postgres.write.dsn"))
			if err != nil {
				log.Fatalf("got an error while initialize database migrations, error: %s", err)
			}

			if err := m.Up(); err != nil {
				log.Fatalf("got an error while execute database migrations, error: %s", err)
			}
		},
	}

	migrationdown = &cobra.Command{
		Use:   "migrate:down",
		Short: "Execute Database Migration Down",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New("file://files/migration", cfg.GetString("postgres.write.dsn"))
			if err != nil {
				log.Fatal(err)
			}
			if err = m.Down(); err != nil {
				log.Fatal(err)
			}
		},
	}
)
