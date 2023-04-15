package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	migrate "github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/linzhengen/mii-go/app/infrastructure/persistence/mysql"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/di"
	"github.com/linzhengen/mii-go/pkg/logger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	envCfg := config.New(ctx)
	db, err := mysql.NewConn(envCfg.MySQL)
	if err != nil {
		logger.Severef("failed connect db, err: %v", err)
	}
	if err := dbMigrate(envCfg, db); err != nil {
		logger.Severef("failed migration, err: %v", err)
	}
	c := di.NewApi(envCfg, db)
	var httpHandler http.Handler
	if err := c.Invoke(func(h http.Handler) {
		httpHandler = h
	}); err != nil {
		logger.Severef("invoke httpHandler err: %v", err)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", envCfg.WebHost, envCfg.WebPort),
		Handler: httpHandler,
	}
	go func() {
		logger.Infof("staring serve, addr: %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Severef("listen: %v", err)
		}
	}()

	<-ctx.Done()

	stop()
	logger.Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Severef("Server forced to shutdown: %v", err)
	}

	logger.Info("Server successfully stopped")
}

//go:embed migrations/mysql/*.sql
var migrationsFs embed.FS

func dbMigrate(envCfg config.EnvConfig, db *sql.DB) error {
	if envCfg.Migration.Auto {
		d, err := iofs.New(migrationsFs, "migrations/mysql")
		if err != nil {
			logger.Severef("failed get migrations", err)
		}
		driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
		if err != nil {
			logger.Severef("failed create mysql migration instance, err: %s", err)
		}
		m, err := migrate.NewWithInstance(
			"iofs",
			d,
			"mysql",
			driver,
		)
		return m.Up()
	}
	return nil
}
