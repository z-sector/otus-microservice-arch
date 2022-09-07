package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"

	"github.com/wuzyk/otus-microservice-arch/05/internal/handlers"
	"github.com/wuzyk/otus-microservice-arch/05/internal/logger"
	"github.com/wuzyk/otus-microservice-arch/05/internal/repo"
	"github.com/wuzyk/otus-microservice-arch/05/internal/server"
)

const (
	httpAddr = ":8000"

	keySize = 2048

	shutdownTimeout   = 5 * time.Second
	readHeaderTimeout = 10 * time.Second
)

func main() {
	zap.ReplaceGlobals(logger.NewZap(zap.InfoLevel))

	conn := pgdriver.NewConnector(
		pgdriver.WithAddr(os.Getenv("PG_ADDR")),
		pgdriver.WithUser(os.Getenv("PG_USR")),
		pgdriver.WithPassword(os.Getenv("PG_PWD")),
		pgdriver.WithDatabase(os.Getenv("PG_DBNAME")),
		pgdriver.WithInsecure(true),
	)
	db := bun.NewDB(sql.OpenDB(conn), pgdialect.New())

	userRepo := repo.NewUserRepo(db)

	key, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		zap.L().Fatal("failed to generate private key")
	}

	authHandler := handlers.NewAuthHandler(userRepo, key)

	router := server.SetupAuthRouter(authHandler)

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
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
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
