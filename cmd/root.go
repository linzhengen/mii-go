package cmd

import (
	"context"
	"database/sql"
	"os/signal"
	"syscall"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Severef("failed execute, err: %v", err)
	}
}

func withInit(f func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB)) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	envCfg := config.New(ctx)
	db, err := mysql.NewConn(envCfg.MySQL)
	if err != nil {
		logger.Severef("failed connect db, err: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Severef("failed close db, err: %v", err)
		}
	}(db)
	f(ctx, stop, envCfg, db)
}
