package cmd

import (
	"context"
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

//go:embed migrations/mysql/*.sql
var migrationsFs embed.FS

var dbMigrateCmd = &cobra.Command{
	Use: "db-migrate",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			d, err := iofs.New(migrationsFs, "migrations/mysql")
			if err != nil {
				logger.Severef("failed get migrations", err)
			}
			driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
			if err != nil {
				logger.Severef("failed create mysql instance driver, err: %s", err)
			}
			m, err := migrate.NewWithInstance(
				"iofs",
				d,
				"mysql",
				driver,
			)
			if err != nil {
				logger.Severef("failed create mysql migration instance, err: %s", err)
			}
			if err := m.Up(); err != nil {
				logger.Severef("failed migrate up, err: %s", err)
			}
		})
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
