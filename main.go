package main

import (
	"context"
	"database/sql"
	"embed"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/linzhengen/mii-go/di"
	"google.golang.org/grpc"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/internal/infrastructure/persistence/mysql"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(
		restCmd,
		grpcCmd,
		grpcGWCmd,
		dbMigrateCmd,
		runCmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		logger.Severef("failed execute, err: %v", err)
	}
}

var rootCmd = &cobra.Command{
	Use: "mii",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var grpcCmd = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db, rootCmd)
			var s *grpc.Server
			if err := c.Invoke(func(server *grpc.Server) {
				s = server
			}); err != nil {
				logger.Severef("invoke grpc server err: %v", err)
			}
			lis, err := net.Listen("tcp", envCfg.Grpc.Addr())
			if err != nil {
				logger.Severef("failed to listen: %v", err)
			}
			logger.Infof("server listening at %v", lis.Addr())
			go func() {
				if err := s.Serve(lis); err != nil {
					logger.Severef("failed to serve: %v", err)
				}
			}()

			// Listen for the interrupt signal.
			<-ctx.Done()

			// Restore default behavior on the interrupt signal and notify user of shutdown.
			stop()
			logger.Info("shutting down gracefully")
			ch := make(chan struct{})
			go func() {
				defer close(ch)
				// close listeners to stop accepting new connections,
				// will block on any existing transports
				s.GracefulStop()
			}()
			select {
			case <-ch:
				logger.Infof("graceful stopped")
			case <-time.After(10 * time.Second):
				// took too long, manually close open transports
				// e.g. watch streams
				logger.Infof("graceful stop timeout, force stop!!")
				s.Stop()
				<-ch
			}
			logger.Info("server successfully stopped")
		})
	},
}

var grpcGWCmd = &cobra.Command{
	Use: "grpc-gw",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db, rootCmd)
			var serveMux *runtime.ServeMux
			if err := c.Invoke(func(mux *runtime.ServeMux) {
				serveMux = mux
			}); err != nil {
				logger.Severef("invoke serveMux err: %v", err)
			}
			srv := &http.Server{
				Addr:    envCfg.GrpcGW.Addr(),
				Handler: serveMux,
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
				logger.Severef("server forced to shutdown: %v", err)
			}

			logger.Info("server successfully stopped")
		})
	},
}

var restCmd = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db, rootCmd)
			var httpHandler http.Handler
			if err := c.Invoke(func(h http.Handler) {
				httpHandler = h
			}); err != nil {
				logger.Severef("invoke httpHandler err: %v", err)
			}
			srv := &http.Server{
				Addr:    envCfg.Rest.Addr(),
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
				logger.Severef("server forced to shutdown: %v", err)
			}

			logger.Info("server successfully stopped")
		})
	},
}

var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db, rootCmd)
			var _ *cobra.Command
			if err := c.Invoke(func(c *cobra.Command) {
				_ = c
			}); err != nil {
				logger.Severef("invoke command err: %v", err)
			}
			<-ctx.Done()
			stop()
		})
	},
}

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
