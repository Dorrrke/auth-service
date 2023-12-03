package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dorrrke/auth-service/internal/app"
	"github.com/Dorrrke/auth-service/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//инициальизировать объект конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//иницаилизация логгера
	log := setupLogger(cfg.Env)
	log.Info("starting application",
		slog.Any("cfg", cfg))

	// TODO: инициализация приложения (app)
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCSrv.MustRun()

	// TODO: запустить gRPC-сервер приложения

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSrv.Stop()

	log.Info("application stop")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
