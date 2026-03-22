package cmd

import (
	"log/slog"
	"os"
)

func Execute() {
	cfg := config{
		port: ":8080",
		db:   databaseConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has faild to start", "error", err)
		os.Exit(1)
	}
}
