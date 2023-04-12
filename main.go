package main

import (
	"context"
	"fmt"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/di"
	"github.com/linzhengen/mii-go/pkg/logger"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	envCfg := config.New(ctx)
	c := di.NewApi(envCfg)
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
		log.Printf("staring serve, addr: %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server successfully stopped")
}
