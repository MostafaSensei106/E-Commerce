package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/MostafaSensei106/E-Commerce/internal/env"
	"github.com/jackc/pgx/v5"
)

func Execute() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Config
	cfg := config{
		port: ":8080",
		db: databaseConfig{
			dsn: env.GetString("Sensei-E-Commerce-DB-DSN", "host=localhost user=root password=root dbname=ecommerce sslmode=disable"),
		},
	}

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dns", cfg.db.dsn)

	// API
	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has faild to start", "error", err)
		os.Exit(1)
	}
}
