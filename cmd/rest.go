package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/di"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db)
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
		})
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
