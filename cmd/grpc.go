package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/di"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcCmd = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		withInit(func(ctx context.Context, stop context.CancelFunc, envCfg config.EnvConfig, db *sql.DB) {
			c := di.NewDI(envCfg, db)
			var s *grpc.Server
			if err := c.Invoke(func(server *grpc.Server) {
				s = server
			}); err != nil {
				logger.Severef("invoke grpc server err: %v", err)
			}
			lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", envCfg.GrpcHost, envCfg.GrpcPort))
			if err != nil {
				logger.Severef("failed to listen: %v", err)
			}
			logger.Infof("server listening at %v", lis.Addr())
			go func() {
				if err := s.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			// Listen for the interrupt signal.
			<-ctx.Done()

			// Restore default behavior on the interrupt signal and notify user of shutdown.
			stop()
			log.Println("shutting down gracefully")
			ch := make(chan struct{})
			go func() {
				defer close(ch)
				// close listeners to stop accepting new connections,
				// will block on any existing transports
				s.GracefulStop()
			}()
			select {
			case <-ch:
				log.Printf("Graceful stopped")
			case <-time.After(10 * time.Second):
				// took too long, manually close open transports
				// e.g. watch streams
				log.Printf("Graceful stop timeout, force stop!!")
				s.Stop()
				<-ch
			}
			log.Println("Server successfully stopped")
		})
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
