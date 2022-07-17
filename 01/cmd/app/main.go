package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wuzyk/otus-microservice-arch/01/internal/handlers"
	"github.com/wuzyk/otus-microservice-arch/01/internal/server"
)

const (
	httpAddr = ":8000"
)

func main() {
	zap.ReplaceGlobals(NewZap(zap.InfoLevel))

	middlwares := &server.Middlewares{}
	handlers := &server.Handlers{
		HealthCheck: handlers.NewHealthCheck(),
	}

	router := server.SetupRouter(middlwares, handlers)

	zap.L().Info("server started")

	initHTTPServer(router, httpAddr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("server stopped")
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

func NewZap(logLevel zapcore.Level) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(config)

	return zap.New(zapcore.NewCore(encoder, os.Stdout, logLevel))
}
