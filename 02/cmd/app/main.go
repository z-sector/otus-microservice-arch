package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/wuzyk/otus-microservice-arch/02/internal/handlers"
	"github.com/wuzyk/otus-microservice-arch/02/internal/logger"
	"github.com/wuzyk/otus-microservice-arch/02/internal/server"
)

const (
	httpAddr = ":8000"

	shutdownTimeout = 5 * time.Second
)

func main() {
	zap.ReplaceGlobals(logger.NewZap(zap.InfoLevel))

	middlwares := &server.Middlewares{}
	handlers := &server.Handlers{
		HealthCheck: handlers.NewHealthCheck(),
	}

	router := server.SetupRouter(middlwares, handlers)

	zap.L().Info("server started")

	server := initHTTPServer(router, httpAddr)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		zap.L().Info("gracefully shutting down...")

		err := server.Shutdown(ctx)
		if err != nil {
			zap.L().Error("failed to shutdown server")
		}

		zap.L().Info("server stopped")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initHTTPServer(router http.Handler, addr string) *http.Server {
	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(
				"http server error",
				zap.Error(err),
				zap.String("addr", addr),
			)
		}
	}()

	return &srv
}
