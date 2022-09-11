package main

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"

	"github.com/wuzyk/otus-microservice-arch/06/internal/logger"
	"github.com/wuzyk/otus-microservice-arch/06/internal/storage/migrations"
)

const (
	migrateTimeout = 30 * time.Second
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

	ctx, cancel := context.WithTimeout(context.Background(), migrateTimeout)
	defer cancel()

	migrator := migrate.NewMigrator(db, migrations.Migrations, migrate.WithMarkAppliedOnSuccess(true))

	err := migrator.Init(ctx)
	if err != nil {
		zap.L().Fatal("failed to create migrations table", zap.Error(err))
		return
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		zap.L().Fatal("failed to run migrations", zap.Error(err))
		return
	}
	if group.IsZero() {
		zap.L().Info("no migrations to run")
	} else {
		zap.L().Info("migrated to " + group.String())
	}
}
