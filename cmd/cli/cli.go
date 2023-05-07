package main

import (
	"context"
	"database/sql"
	"os/signal"
	"syscall"

	"github.com/linzhengen/mii-go/di"
	"github.com/linzhengen/mii-go/internal/interface/cmd/register"
	"github.com/spf13/cobra"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql"
	"github.com/linzhengen/mii-go/pkg/logger"
)

func main() {
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

	var rootCmd = &cobra.Command{
		Use:   "mii",
		Short: "mii is a Golang project template",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	c := di.NewDI(envCfg, db)
	var commands register.Commands
	if err := c.Invoke(func(cmds register.Commands) {
		commands = cmds
	}); err != nil {
		logger.Severef("invoke commands err: %v", err)
	}
	rootCmd.SetContext(ctx)
	rootCmd.AddCommand(commands...)
}
