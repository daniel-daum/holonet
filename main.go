package main

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func initalize() {

	devLogger := tint.NewTextHandler(os.Stderr, &tint.Options{Level: slog.LevelDebug})
	// jsonHandler := slog.NewJSONHandler(os.Stderr, nil)

	aslog := slog.New(devLogger)

	slog.SetDefault(aslog)

}

func main() {

	config, err := loadEnv()

	if err != nil {
		slog.Error("Invalid config.", "error", err)
		os.Exit(1)
	}

	initalize()

	slog.Info("starting Holonet Server", "env", config.Env(), "port", config.Port())
	slog.Warn("warning")
	slog.Error("error")
	slog.Debug("debug")
}
